package proto

//easyjson:json
type InitMessage struct {
	Address string `json:"address"`
}

//easyjson:json
type InitReply struct {
	Challenge string `json:"challenge"`
	Signature string `json:"signature"`
	Error     string `json:"error"`
}
