package lsp

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type TextDocumentItem struct {
	TextDocumentIdentifier
	LanguageId string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

type VersionTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type TextDocumentSyncOptions struct {
	OpenClose bool `json:"openClose"`
	WillSave  bool `json:"willSave"`
	Change    int  `json:"change"`
}
