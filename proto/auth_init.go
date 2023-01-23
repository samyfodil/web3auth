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

func MarshaledInitReplyError(err error) []byte {
	bytes, _ := InitReply{Error: err.Error()}.MarshalJSON()
	return bytes
}
