package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func validateNewUser(t *testing.T, user *User) {
	assert.Equal(t, "first", user.FirstName)
	assert.Equal(t, "last", user.LastName)
	assert.Equal(t, "user@crazy.net", user.Email)
}

func TestCreateNew(t *testing.T) {
	user, err := NewUser("first", "last", "user@crazy.net")
	if assert.Nil(t, err) {
		validateNewUser(t, user)

		events := eventStore.GetEvents()
		assert.Equal(t, 1, len(events))

		assert.NotEqual(t, "", user.AggregateId)
	}
}

func TestNewFromEvents(t *testing.T) {
	createEvent := UserCreated{
		FirstName: "first",
		LastName:  "last",
		Email:     "user@crazy.net",
	}

	user, err := NewUserFromHistory([]interface{}{createEvent})
	if assert.Nil(t, err) {
		validateNewUser(t, user)
	}

}
