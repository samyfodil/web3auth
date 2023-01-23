package proto

//easyjson:json
type Message struct {
	Address string `json:"address"`
	Init    struct {
		Challenge string `json:"challenge"`
		Signature string `json:"signature"`
	} `json:"init"`
	Signature string `json:"signature"`
}

//easyjson:json
type Response struct {
	Token string `json:"token"`
	Error string `json:"error"`
}
