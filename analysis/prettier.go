package analysis

import (
	"bytes"
	"os/exec"
)

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
