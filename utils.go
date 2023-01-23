package web3auth

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func parseSignature(signature string) ([]byte, error) {
	signHex := []byte(signature)

	if len(signHex) >= 2 && (signHex[0] == '0' && (signHex[1] == 'x' || signHex[1] == 'X')) {
		signHex = signHex[2:]
	}

	sign := make([]byte, hex.DecodedLen(len(signHex)))
	_, err := hex.Decode(sign, signHex)
	if err != nil {
		return nil, err
	}

	sign[64] -= 27

	return sign, nil
}

func walletAddrFromPubKey(pkey []byte) string {
	var buf []byte
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pkey[1:]) // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]
	return "0x" + hex.EncodeToString(address)
}
