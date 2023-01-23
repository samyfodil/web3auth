package web3auth

import "github.com/samyfodil/web3auth/proto"

func MarshaledReplyError(err error) []byte {
	res := &proto.Response{Error: err.Error()}
	bytes, _ := res.MarshalJSON()
	return bytes
}

func MarshaledInitReplyError(err error) []byte {
	res := &proto.InitReply{Error: err.Error()}
	bytes, _ := res.MarshalJSON()
	return bytes
}
