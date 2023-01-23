package key

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKey(t *testing.T) {
	assert := assert.New(t)

	key, err := New(big.NewInt(1).Bytes())
	assert.Nil(err)
	assert.NotNil(key)
}

func TestKeyBytes(t *testing.T) {
	assert := assert.New(t)

	keyBytes := make([]byte, 32)
	big.NewInt(1).FillBytes(keyBytes)
	key, err := New(keyBytes)

	assert.Nil(err)
	assert.NotNil(key)

	bytes := key.Bytes()
	assert.NotNil(bytes)
	assert.Equal(32, len(bytes))
	assert.Equal(big.NewInt(1), big.NewInt(0).SetBytes(bytes))
}

func TestKeyPublic(t *testing.T) {
	assert := assert.New(t)

	keyBytes := make([]byte, 32)
	big.NewInt(1).FillBytes(keyBytes)
	key, err := New(keyBytes)

	assert.Nil(err)
	assert.NotNil(key)

	pubkey, err := key.Public()
	assert.Nil(err)
	assert.NotNil(pubkey)
}
