package web3auth

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshaledReplyError(t *testing.T) {
	assert := assert.New(t)

	err := errors.New("test error")
	assert.Equal(`{"error":"test error"}`, string(MarshaledReplyError(err)))
}

func TestMarshaledInitReplyError(t *testing.T) {
	assert := assert.New(t)

	err := errors.New("test error")
	assert.Equal(`{"error":"test error"}`, string(MarshaledInitReplyError(err)))
}
