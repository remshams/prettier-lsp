package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

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
	case "textDocument/willSaveWaitUntil":
		var request lsp.WillSaveWaitUntilTextDocumentNotification
		err := json.Unmarshal(content, &request)
		if err != nil {
			logger.Printf("Could not parse request: %s", err)
		}
		logger.Printf("Will save file %s", request.Params.TextDocument.URI)
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
