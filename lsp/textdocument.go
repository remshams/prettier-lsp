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
	OpenClose         bool `json:"openClose"`
	WillSaveWaitUntil bool `json:"willSaveWaitUntil"`
	Change            int  `json:"change"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type TextEditResponse struct {
	Range   `json:"range"`
	NewText string `json:"newText"`
}
