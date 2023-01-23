package key

import (
	"errors"
	"strings"

	ethereum "github.com/taubyte/go-sdk/ethereum/client"
)

type _key struct {
	bytes []byte
}

func New[T []byte | string](key T) (Key, error) {
	iface := any(key)
	switch iface.(type) {
	case []byte:
		bytes := iface.([]byte)
		k := &_key{
			bytes: make([]byte, len(bytes)),
		}
		copy(k.bytes, bytes)
		return k, nil
	case string:
		hexkey := strings.ToLower(iface.(string))
		if strings.HasPrefix(hexkey, "0x") == true {
			hexkey = hexkey[2:]
		}
		bytes, err := ethereum.HexToECDSABytes(hexkey)
		if err != nil {
			return nil, err
		}
		return &_key{bytes: bytes}, nil
	}
	return nil, errors.New("Should never happen!")
}

func (k *_key) Bytes() []byte {
	return k.bytes
}

func (k *_key) Public() (Key, error) {
	bytes, err := ethereum.PublicKeyFromPrivate(k.bytes)
	if err != nil {
		return nil, err
	}

	return New(bytes)
}
