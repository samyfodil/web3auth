package web3auth

import (
	"testing"
	"time"

	"github.com/samyfodil/web3auth/key"
)

func TestNew(t *testing.T) {
	if New() != instantiatorSingleton {
		t.Error("expected instantiatorSingleton")
	}
}

func TestNewVerifier(t *testing.T) {
	k, _ := key.New([]byte("0x0000000000000000000000000000000000000000000000000000000000000001"))
	verifier := New().Verifier(k)
	if verifier == nil {
		t.Error("expected verifier to not be nil")
	}
}

func TestNewIssuer(t *testing.T) {
	k, _ := key.New([]byte("0x0000000000000000000000000000000000000000000000000000000000000001"))
	issuer := New().Issuer(k, time.Second)
	if issuer == nil {
		t.Error("expected issuer to not be nil")
	}
}
