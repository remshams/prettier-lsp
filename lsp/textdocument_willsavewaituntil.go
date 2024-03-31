package lsp

type WillSaveWaitUntilTextDocumentNotification struct {
	Notification
	Params WillSaveWaitUntilTextDocumentParams `json:"params"`
}

type WillSaveWaitUntilTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type TextEditResponse struct {
	Range
	NewText string `json:"newText"`
}

type WillSaveWaitUnitlTextDocumentResponse struct {
	Response
	Result []TextEditResponse `json:"result"`
}

func CreateWillSaveWaitUntilTextDocumentResponse(id int, startPosition Position, text string) WillSaveWaitUnitlTextDocumentResponse {
	endPosition := Position{
		Line:      startPosition.Line,
		Character: startPosition.Character + len(text),
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
				NewText: text,
			},
		},
	}
}
