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

		events, err := myEventStore.GetEvents(user.AggregateId)
		if assert.Nil(t, err) {
			assert.Equal(t, 1, len(events))
			assert.NotEqual(t, "", user.AggregateId)
		}
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

func TestUpdateFirstName(t *testing.T) {
	user, err := NewUser("first", "last", "user@crazy.net")
	if assert.Nil(t, err) {
		err := user.UpdateFirstName("foo")
		assert.Nil(t, err)
		assert.Equal(t, "foo", user.FirstName)

		events, err := myEventStore.GetEvents(user.AggregateId)
		if assert.Nil(t, err) {
			assert.Equal(t, 2, len(events))

			userCopy, err := NewUserFromHistory(events)
			assert.Nil(t, err)
			assert.Equal(t, user, userCopy)
		}
	}
}
