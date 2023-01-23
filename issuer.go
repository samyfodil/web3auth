package web3auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/samyfodil/web3auth/key"
	"github.com/samyfodil/web3auth/proto"
	vmRand "github.com/taubyte/go-sdk/crypto/rand"
	ethereum "github.com/taubyte/go-sdk/ethereum/client"
	"github.com/taubyte/utils/multihash"
)

type issuer struct {
	sk  key.Key
	pk  key.Key
	ttl int64
	err error
}

func newIssuer(sk key.Key, ttl time.Duration) Issuer {
	pk, err := sk.Public()
	return &issuer{
		sk:  sk,
		pk:  pk,
		ttl: int64(ttl / time.Second),
		err: err,
	}
}

// Given an Init Message it will issue a challenge
// the challenge is verifiable by any node in the network given they
// have the public key of the issuer
func (i *issuer) Challenge(message *proto.InitMessage) *proto.InitReply {
	if i.err != nil {
		return &proto.InitReply{
			Error: i.err.Error(),
		}
	}

	nonce, err := rand.Int(vmRand.NewReader(), big.NewInt(math.MaxInt64))
	if err != nil {
		return &proto.InitReply{
			Error: err.Error(),
		}
	}

	challenge := multihash.Hash(message.Address + nonce.String())

	sign, err := ethereum.SignMessage([]byte(challenge), i.sk.Bytes())
	if err != nil {
		return &proto.InitReply{
			Error: err.Error(),
		}
	}

	sign[64] += 27

	return &proto.InitReply{
		Challenge: challenge,
		Signature: "0x" + hex.EncodeToString(sign),
	}
}

func (i *issuer) Issue(message *proto.Message) *proto.Response {
	if i.err != nil {
		return &proto.Response{
			Error: i.err.Error(),
		}
	}

	err := i.verifyChallengeIssuerSignature(message)
	if err != nil {
		return &proto.Response{
			Error: i.err.Error(),
		}
	}

	err = i.verifyChallengeWalletSignature(message)
	if err != nil {
		return &proto.Response{
			Error: i.err.Error(),
		}
	}

	token, err := i.issueToken(message)
	if err != nil {
		return &proto.Response{
			Error: i.err.Error(),
		}
	}

	return &proto.Response{
		Token: token,
	}
}

func (i *issuer) verifyChallengeIssuerSignature(message *proto.Message) error {
	issuerSig, err := parseSignature(message.Init.Signature)
	if err != nil {
		return err
	}

	challenge := []byte(message.Init.Challenge)
	err = ethereum.VerifySignature(challenge, i.pk.Bytes(), issuerSig)
	if err != nil {
		return err
	}

	return nil
}

func (i *issuer) verifyChallengeWalletSignature(message *proto.Message) error {
	challenge := []byte("\x19Ethereum Signed Message:\n" +
		strconv.Itoa(len(message.Init.Challenge)) +
		message.Init.Challenge)

	walletSig, err := parseSignature(message.Signature)
	if err != nil {
		return err
	}

	pub, err := ethereum.PublicKeyFromSignedMessage(challenge, walletSig)
	if err != nil {
		return err
	}

	err = ethereum.VerifySignature(challenge, pub, walletSig)
	if err != nil {
		return err
	}

	waddr := walletAddrFromPubKey(pub)
	if strings.ToLower(waddr) != strings.ToLower(message.Address) {
		return errors.New("Wallet signature does not match address: " + waddr + "!=" + message.Address)
	}

	return nil
}

func (i *issuer) issueToken(message *proto.Message) (string, error) {
	waddr := strings.ToLower(message.Address)
	expire := time.Now().Unix() + i.ttl
	tokenInfo := "addr=" + waddr + ";expire=" + strconv.FormatInt(expire, 10)
	sign, err := ethereum.SignMessage(
		[]byte("\x19Ethereum Signed Message:\n"+strconv.Itoa(len(tokenInfo))+tokenInfo),
		i.sk.Bytes(),
	)
	if err != nil {
		return "", err
	}

	sign[64] += 27

	token := []byte(tokenInfo + ";sign=0x" + hex.EncodeToString(sign))

	return base64.StdEncoding.EncodeToString(token), nil
}
