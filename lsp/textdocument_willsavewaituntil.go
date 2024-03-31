package lsp

import (
	"strings"
)

type WillSaveWaitUntilTextDocumentNotification struct {
	Request
	Params WillSaveWaitUntilTextDocumentParams `json:"params"`
}

type WillSaveWaitUntilTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type TextEditResponse struct {
	Range   `json:"range"`
	NewText string `json:"newText"`
}

type WillSaveWaitUnitlTextDocumentResponse struct {
	Response
	Result []TextEditResponse `json:"result"`
}

func CreateWillSaveWaitUntilTextDocumentResponse(id int, oldText string, newText string) WillSaveWaitUnitlTextDocumentResponse {
	lines := strings.Split(oldText, "\n")
	lastLineNumber := len(lines) - 1
	lastCharacterNumber := len(lines[lastLineNumber])
	startPosition := Position{
		Line:      0,
		Character: 0,
	}
	endPosition := Position{
		Line:      lastLineNumber,
		Character: lastCharacterNumber,
	}
	return WillSaveWaitUnitlTextDocumentResponse{
		Response: Response{
			RPC: "2.0",
			Id:  &id,
		},
		Result: []TextEditResponse{
			{
				Range: Range{
					Start: startPosition,
					End:   endPosition,
				},
				NewText: newText,
			},
		},
	}
}
