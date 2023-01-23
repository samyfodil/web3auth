package web3auth

import (
	"time"

	"github.com/samyfodil/web3auth/key"
)

type instantiator struct{}

var instantiatorSingleton = &instantiator{}

func New() Instantiator {
	return instantiatorSingleton
}

func (i *instantiator) Issuer(sk key.Key, ttl time.Duration) Issuer {
	return newIssuer(sk, ttl)
}

func (i *instantiator) Verifier(pk key.Key) Verifier {
	return newVerifier(pk)
}
