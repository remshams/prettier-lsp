package lsp

type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionTextDocumentIdentifier    `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}
