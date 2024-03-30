package main

import (
	"bufio"
	"log"
	"os"

	"github.com/remshams/prettier-lsp/rpc"
)

func main() {
	logger := getLogger("/home/remshams/coding/prettier-lsp/prettier-lsp.log")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println("Could not decode message")
			continue
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("Method: %s", method)
	logger.Printf("Content %s", content)
}

func getLogger(fileName string) *log.Logger {
	logfile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Could not setup logger")
	}
	return log.New(logfile, "[prettier-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
