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
	TextDocumentSync TextDocumentSyncOptions `json:"textDocumentSync"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func CreateInitializeResult(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			Id:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: TextDocumentSyncOptions{
					OpenClose: true,
					WillSave:  true,
					Change:    1,
				},
			},
			ServerInfo: ServerInfo{
				Name:    "prettier-lsp",
				Version: "0.0.1",
			},
		},
	}
}
