package web3auth

import "github.com/samyfodil/web3auth/proto"

func MarshaledReplyError(err error) []byte {
	bytes, _ := proto.Response{Error: err.Error()}.MarshalJSON()
	return bytes
}

func MarshaledInitReplyError(err error) []byte {
	bytes, _ := proto.InitReply{Error: err.Error()}.MarshalJSON()
	return bytes
}
