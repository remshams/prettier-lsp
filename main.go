package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/remshams/prettier-lsp/.git/lsp"
	"github.com/remshams/prettier-lsp/rpc"
)

func main() {
	logger := getLogger("/home/remshams/coding/prettier-lsp/prettier-lsp.log")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println("Could not decode message")
			continue
		}
		handleMessage(logger, writer, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, method string, content []byte) {
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
