package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	Id     int    `json:"id"`
	Method string `json:"method"`
}

type Response struct {
	RPC string `json:"jsonrpc"`
	Id  *int   `json:"id"`
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}

func CreateResponse(id int) Response {
	return Response{
		RPC: "2.0",
		Id:  &id,
	}
}
