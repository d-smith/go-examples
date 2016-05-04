package aggregates

import (
	"github.com/d-smith/go-examples/es2"
	"github.com/d-smith/go-examples/es2/sample"
	. "github.com/lsegal/gucumber"
	"github.com/stretchr/testify/assert"
)

func init() {
	var eventHistory []es2.Event
	var userFromHistory *sample.User

	Given(`^an event sourced aggregate's event history$`, func() {
		eventHistory = []es2.Event{
			es2.Event{
				Payload: sample.UserCreated{
					AggregateId: "123",
					FirstName:   "orig first",
					LastName:    "orig last",
					Email:       "orig@email.com",
				},
			},
			es2.Event{
				Payload: sample.UserFirstNameUpdated{
					NewFirst: "new first",
				},
			},
		}

	})

	When(`^I instantiate the aggregate from its history$`, func() {
		userFromHistory = sample.NewUserFromHistory(eventHistory)
		assert.NotNil(T, userFromHistory, "Nil user build from history")
	})

	Then(`^the instance state is correct$`, func() {
		assert.Equal(T, "123", userFromHistory.ID)
		assert.Equal(T, "new first", userFromHistory.FirstName)
		assert.Equal(T, "orig last", userFromHistory.LastName)
		assert.Equal(T, "orig@email.com", userFromHistory.Email)
	})

	And(`^there are no uncommitted events$`, func() {
		assert.Equal(T, 0, len(userFromHistory.Events))
	})

}
