package rpc

import (
	"bytes"
	"fmt"
	"testing"
)

var content = "{\"method\":\"hello\"}"

func createMessage(content string) (string, []byte, int) {
	contentLength := len(content)
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", contentLength)
	message := header + content
	messageBytes := []byte(message)
	return message, messageBytes, contentLength
}

func TestSplitIncomplete(t *testing.T) {
	_, messageBytes, _ := createMessage(content)
	before, _, _ := bytes.Cut(messageBytes, []byte(" "))
	advance, _, err := Split(before, false)
	if err != nil {
		t.Fatalf("Could not split")
	}
	if advance != 0 {
		t.Fatalf("Bytes are incomplete, should return 0 for advance, got: %d", advance)
	}
}

func TestSplitJustContentType(t *testing.T) {
	_, messageBytes, _ := createMessage(content)
	before, _, _ := bytes.Cut(messageBytes, []byte{'\r', '\n', '\r', '\n'})
	advance, _, err := Split(before, false)
	if err != nil {
		t.Fatalf("Could not split")
	}
	if advance != 0 {
		t.Fatalf("Only content type but not content, should return 0 for advance, got: %d", advance)
	}
}

func TestSplitContentIncomplete(t *testing.T) {
	_, messageBytes, _ := createMessage(content)
	before, _, _ := bytes.Cut(messageBytes, []byte("method"))
	advance, _, err := Split(before, false)
	if err != nil {
		t.Fatalf("Could not split")
	}
	if advance != 0 {
		t.Fatalf("Content incomplete, should return 0 for advance, got %d", advance)
	}
}

func TestSplitCompleteContent(t *testing.T) {
	message, messageBytes, contentLength := createMessage(content)
	advance, data, err := Split(messageBytes, false)
	if err != nil {
		t.Fatalf("Could not split")
	}

	// "Content-Length: {contentLength}" + "\r\n\r\n" + "Hello World"
	totalLength := 18 + 4 + contentLength
	if advance != totalLength {
		t.Fatalf("Content complete, should advance by %d, got %d", totalLength, advance)
	}
	if string(data) != message {
		t.Fatalf("Content complete, should return all data, got \"%s\"", data)
	}
}

func TestDecodeMessage(t *testing.T) {
	_, messageBytes, _ := createMessage(content)
	decoded, _, err := DecodeMessage(messageBytes)
	if err != nil {
		t.Fatalf("Could not decode message: %s", err)
	}
	if decoded != "hello" {
		t.Fatalf("Decoded content does not match content, got \"%s\"", decoded)
	}
}

func TestDecodeNotABaseMessage(t *testing.T) {
	_, messageBytes, _ := createMessage("Hello World")
	_, _, err := DecodeMessage(messageBytes)
	if err == nil {
		t.Fatalf("Invalid content, should not decode message")
	}
}
