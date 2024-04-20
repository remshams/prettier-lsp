package lsp

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeParamsRequest struct {
	ClientInfo `json:"clientInfo"`
}

type InitializeRequest struct {
	Request
	Params InitializeParamsRequest `json:"params"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync           TextDocumentSyncOptions `json:"textDocumentSync"`
	DocumentFormattingProvider bool                    `json:"documentFormattingProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func CreateInitializeResult(id int, name string, version string) InitializeResponse {
	return InitializeResponse{
		Response: CreateResponse(id),
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: TextDocumentSyncOptions{
					OpenClose:         true,
					WillSaveWaitUntil: false,
					Change:            1,
				},
				DocumentFormattingProvider: true,
			},
			ServerInfo: ServerInfo{
				Name:    name,
				Version: version,
			},
		},
	}
}
