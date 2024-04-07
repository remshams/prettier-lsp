package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/remshams/prettier-lsp/analysis"
	"github.com/remshams/prettier-lsp/lsp"
	"github.com/remshams/prettier-lsp/rpc"
)

func main() {
	logger := getLogger("/home/remshams/coding/prettier-lsp/prettier-lsp.log")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	state := analysis.NewState(logger)

	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println("Could not decode message")
			continue
		}
		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, content []byte) {
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		err := json.Unmarshal(content, &request)
		if err != nil {
			logger.Printf("Could not parsed request: %s", err)
		}
		logger.Printf("Connected to %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
		msg := lsp.CreateInitializeResult(request.Id)
		writeResponse(writer, msg)
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		err := json.Unmarshal(content, &request)
		if err != nil {
			logger.Printf("Could not parse request: %s", err)
		}
		logger.Printf("Opened file %s", request.Params.TextDocument.URI)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		err := json.Unmarshal(content, &request)
		if err != nil {
			logger.Printf("Could not parse request: %s", err)
		}
		for _, change := range request.Params.ContentChanges {
			logger.Printf("Changed file %s", request.Params.TextDocument.URI)
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/formatting":
		var request lsp.FormattingTextDocumentRequest
		err := json.Unmarshal(content, &request)
		if err != nil {
			logger.Printf("Could not parse request: %s", err)
		}
		logger.Printf("Prettier formatting for: %s", request.Params.TextDocument.URI)
		cmd := exec.Command("prettierd", request.Params.TextDocument.URI)

		// Set up input and output buffers
		oldText := state.Documents[request.Params.TextDocument.URI]
		in := bytes.NewBufferString(oldText)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdin = in
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		formattedText := oldText
		err = cmd.Run()
		if err == nil {
			formattedText = out.String()
		} else {
			logger.Printf("Prettier error: %v", err)
			logger.Printf("stderr: %s", stderr.String())
		}
		response := lsp.CreateFormattingTextDocumentResponse(request.Id, state.Documents[request.Params.TextDocument.URI], formattedText)
		writeResponse(writer, response)
	default:
		logger.Printf("Method: %s", method)
	}
}

func getLogger(fileName string) *log.Logger {
	logfile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Could not setup logger")
	}
	return log.New(logfile, "[prettier-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}
