package lsp

import (
	"github.com/remshams/prettier-lsp/analysis"
)

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
	textMetadata := analysis.GetTextMetadata(oldText)
	startPosition := Position{
		Line:      0,
		Character: 0,
	}
	endPosition := Position{
		Line:      textMetadata.LastLineNumber,
		Character: textMetadata.LastLineCharacterNumber,
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
