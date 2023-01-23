[![Release](https://img.shields.io/github/v/release/samyfodil/web3auth.svg)](https://github.com/samyfodil/web3auth/releases)
[![License](https://img.shields.io/github/license/samyfodil/web3auth)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/samyfodil/web3auth)](https://goreportcard.com/report/github.com/samyfodil/web3auth)
[![GoDoc](https://godoc.org/github.com/github.com/samyfodil/web3auth?status.svg)](https://pkg.go.dev/github.com/github.com/samyfodil/web3auth)

# web3auth
This package makes implementing a Web3 Wallets Login possible in few lines of code!


## Usage 
The protocol requires two endpoints to issue a token. The first call provides the client with a verifiable random challenge. While the second call verifies the challenge signature then issues a token that can then be used either on or off-chain.

## Import
``` go
import (
    "github.com/samyfodil/web3auth"
    "github.com/samyfodil/web3auth/key"
)
```


## Issue tokens

### Initialize your issuer
```go

var (
	issuer web3auth.Issuer
)

func init() {
	sk, _ := key.New("HEX-OF-PRIVATE-KEY")
	issuer = web3auth.New().Issuer(
		sk,
		7*24*time.Hour,
	)
}
```


### First endpoint (i.e. /auth/init)
This call will generate a challenge. In this example I'm using a [Taubyte dFunc](https://tau.how).
``` go
//export auth_wallet_init
func authWalletInit(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	defer func() {
		if err != nil {
			h.Write(web3auth.MarshaledInitReplyError(err))
		}
	}()

	body, err := io.ReadAll(h.Body())
	if err != nil {
		return 1
	}

	msg := &proto.InitMessage{}
	err = msg.UnmarshalJSON(body)
	if err != nil {
		return 1
	}

	res := issuer0.Challenge(msg)

	resBytes, err := res.MarshalJSON()
	if err != nil {
		return 1
	}

	h.Write(resBytes)
	h.Return(200)

	return 0
}
```


### Second endpoint (i.e. /auth)
This call verifies the wallet.
```go
//export auth_wallet
func authWallet(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	defer func() {
		if err != nil {
			h.Write(web3auth.MarshaledReplyError(err))
		}
	}()

	body, err := io.ReadAll(h.Body())
	if err != nil {
		return 1
	}

	msg := &proto.Message{}
	err = msg.UnmarshalJSON(body)
	if err != nil {
		return 1
	}

	res := issuer1.Issue(msg)

	resData, _ := res.MarshalJSON()

	h.Write(resData)
	return 0
}
```

## Verify tokens
### Initialize your verifier
```go
var (
	verifier web3auth.Verifier
)

func init() {
	pk, _ := key.New("HEX-OF-PUBLIC-KEY")
	verifier = web3auth.New().Verifier(pk)
}
```


### Verify a token

```go
err := verifier.Validate("token")
```

# Maintainers
 - Samy Fodil @samyfodil