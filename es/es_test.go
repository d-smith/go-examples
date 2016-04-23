package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNew(t *testing.T) {
	user, err := NewUser("first", "last", "user@crazy.net")
	assert.Nil(t, err)
	assert.Equal(t, "first", user.FirstName)
	assert.Equal(t, "first", user.FirstName)
	assert.Equal(t, "user@crazy.net", user.Email)
}
