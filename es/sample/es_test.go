package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
"github.com/d-smith/go-examples/es"
	"time"
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

		//Nothing in the event store yet
		events, err := myEventStore.GetEvents(user.AggregateId)
		assert.NotNil(t,err)

		//Flush to the event stream
		user.Flush(myEventStream)

		//Sleep to let event stream pick up flused event
		time.Sleep(10 * time.Millisecond)

		events, err = myEventStore.GetEvents(user.AggregateId)
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

	user, err := NewUserFromHistory(
		[]es.Event{
			es.Event {
				AggregateId:"123",
				Payload:createEvent,
			},
		})
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

		//Flush to the event stream
		user.Flush(myEventStream)

		//Sleep to let event stream pick up flused event
		time.Sleep(10 * time.Millisecond)

		events, err := myEventStore.GetEvents(user.AggregateId)
		if assert.Nil(t, err) {
			assert.Equal(t, 2, len(events))

			userCopy, err := NewUserFromHistory(events)
			assert.Nil(t, err)
			assert.Equal(t, user.AggregateId, userCopy.AggregateId)
			assert.Equal(t, user.FirstName, userCopy.FirstName)
			assert.Equal(t, user.LastName, userCopy.LastName)
			assert.Equal(t, user.Email, userCopy.Email)
		}
	}
}
