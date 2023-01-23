package web3auth

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/samyfodil/web3auth/key"
	ethereum "github.com/taubyte/go-sdk/ethereum/client"
)

type verifier struct {
	pk key.Key
}

func newVerifier(pk key.Key) Verifier {
	return &verifier{
		pk: pk,
	}
}

func (v *verifier) Validate(token string) error {
	bytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return err
	}

	token = string(bytes)

	tokenElms := make(map[string]string)
	for _, elm := range strings.Split(token, ";") {
		tkns := strings.Split(elm, "=")
		if len(tkns) != 2 {
			continue
		}
		attr := tkns[0]
		value := tkns[1]
		tokenElms[attr] = value
	}

	tokenInfo := "addr=" + tokenElms["addr"] + ";expire=" + tokenElms["expire"]
	message := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(tokenInfo)) + tokenInfo)

	sig, err := parseSignature(tokenElms["sign"])
	if err != nil {
		return err
	}

	err = ethereum.VerifySignature(message, v.pk.Bytes(), sig)
	if err != nil {
		return err
	}

	// parse expire
	expire, err := strconv.ParseInt(tokenElms["expire"], 10, 64)
	if err != nil {
		return err
	}

	if expire < time.Now().Unix() {
		return errors.New("Token expired")
	}

	return nil
}
