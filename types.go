package web3auth

import (
	"time"

	"github.com/samyfodil/web3auth/key"
	"github.com/samyfodil/web3auth/proto"
)

type Instantiator interface {
	Issuer(sk key.Key, ttl time.Duration) Issuer
	Verifier(pk key.Key) Verifier
}

type Issuer interface {
	Challenge(message *proto.InitMessage) *proto.InitReply
	Issue(message *proto.Message) *proto.Response
}
type Verifier interface {
	Validate(token string) error
}
