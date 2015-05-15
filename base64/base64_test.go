package main

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	s := "This is a test of base64 encoding, which is pretty easy"
	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	assert.Nil(t, err)
	assert.Equal(t, s, decoded)
}
