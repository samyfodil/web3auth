package web3auth

import "testing"

func TestWalletAddrFromPubKey(t *testing.T) {
	pkey := []byte{4, 39, 26, 37, 31, 48, 7, 206, 141, 93, 67, 0, 12, 248, 215, 93, 167, 130, 170, 27, 253, 208, 187, 253, 196, 147, 189, 137, 88, 171, 19, 248, 243, 11, 30, 116, 193, 3, 33, 78, 210, 96, 164, 166, 29, 14, 65, 197, 245, 52, 216, 96, 34, 153, 199, 65, 182, 36, 224, 248, 203, 111, 13, 47, 116, 10}
	addr := walletAddrFromPubKey(pkey)
	if addr != "0x2aB2B6dA7aB9B6bAe6aAe6B2B6dA7aB9B6bAe6aA" {
		t.Error("walletAddrFromPubKey failed")
	}
}

func TestParseSignature(t *testing.T) {
	//TODO
}
