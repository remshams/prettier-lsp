package lsp

import "strings"

type FormattingTextDocumentRequest struct {
	Request
	Params FormattingTextDocumentParams `json:"params"`
}

type FormattingTextDocumentParams struct {
	Request
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type FormattingTextDocumentResponse struct {
	Response
	Result []TextEditResponse `json:"result"`
}

func CreateFormattingTextDocumentResponse(id int, oldText string, newText string) FormattingTextDocumentResponse {
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
	return FormattingTextDocumentResponse{
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
