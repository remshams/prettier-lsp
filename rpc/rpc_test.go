package rpc

import (
	"bytes"
	"testing"
)

var content = "Content-Length: 11\r\n\r\nHello World"

func TestSplitIncomplete(t *testing.T) {
	before, _, _ := bytes.Cut([]byte(content), []byte(" "))
	advance, _, err := Split(before, false)
	if err != nil {
		t.Fatalf("Could not split")
	}
	if advance != 0 {
		t.Fatalf("Bytes are incomplete, should return 0 for advance, got: %d", advance)
	}
}

func TestSplitJustContentType(t *testing.T) {
	before, _, _ := bytes.Cut([]byte(content), []byte{'\r', '\n', '\r', '\n'})
	advance, _, err := Split(before, false)
	if err != nil {
		t.Fatalf("Could not split")
	}
	if advance != 0 {
		t.Fatalf("Only content type but not content, should return 0 for advance, got: %d", advance)
	}
}

func TestSplitContentIncomplete(t *testing.T) {
	before, _, _ := bytes.Cut([]byte(content), []byte("World"))
	advance, _, err := Split(before, false)
	if err != nil {
		t.Fatalf("Could not split")
	}
	if advance != 0 {
		t.Fatalf("Content incomplete, should return 0 for advance, got %d", advance)
	}
}

func TestSplitCompleteContent(t *testing.T) {
	advance, data, err := Split([]byte(content), false)
	if err != nil {
		t.Fatalf("Could not split")
	}

	// "Content-Length: 11" + "\r\n\r\n" + "Hello World"
	totalLength := 18 + 4 + 11
	if advance != totalLength {
		t.Fatalf("Content complete, should advance by %d, got %d", totalLength, advance)
	}
	if string(data) != content {
		t.Fatalf("Content complete, should return all data, got \"%s\"", data)
	}
}
