package analysis

import (
	"bytes"
	"os/exec"
	"strings"
)

type TextMetadata struct {
	LastLineNumber          int
	LastLineCharacterNumber int
}

func GetTextMetadata(text string) TextMetadata {
	lines := strings.Split(text, "\n")
	lastLineNumber := len(lines) - 1
	lastCharacterNumber := len(lines[lastLineNumber])
	return TextMetadata{
		LastLineNumber:          lastLineNumber,
		LastLineCharacterNumber: lastCharacterNumber,
	}
}

func FormatWithPrettier(text string, filePath string) (string, error) {
	cmd := exec.Command("prettierd", filePath)
	// Set up input and output buffers
	in := bytes.NewBufferString(text)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdin = in
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
