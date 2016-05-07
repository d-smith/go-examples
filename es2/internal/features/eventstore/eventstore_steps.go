package eventstore

import (
	. "github.com/lsegal/gucumber"
	"github.com/d-smith/go-examples/es2/sample"
	"github.com/d-smith/go-examples/es2/eventstore"
	"github.com/d-smith/go-examples/es2"
	"github.com/stretchr/testify/assert"
)

func init() {
	var user, user2 *sample.User
	var eventStore es2.EventStore


	Given(`^an aggregate$`, func() {
		var err error
		user,err = sample.NewUser("first","last","email")
		assert.Nil(T,err)
	})

	And(`^an event store$`, func() {
		eventStore = eventstore.NewInMemoryEventStore()
	})

	When(`^the aggregate has uncommitted events$`, func() {

	})

	And(`^the events are stored$`, func() {
		err := user.Store(eventStore)
		assert.Nil(T,err)
		assert.Equal(T, 0, len(user.Events))
	})

	Then(`^the events for that aggregate can be retrieved$`, func() {
		eventSets, err := eventStore.RetrieveEvents(user.ID)
		assert.Nil(T,err)
		assert.Equal(T, 1, len(eventSets), "Expected one event set to be retrieved")
	})

	And(`^the aggregate state can be recreated using the events$`, func() {
		eventSets, err := eventStore.RetrieveEvents(user.ID)
		assert.Nil(T,err)
		retUser := sample.NewUserFromHistory(eventSets[0])
		assert.NotNil(T, retUser)
		assert.Equal(T, user.FirstName, retUser.FirstName)
		assert.Equal(T, user.LastName, retUser.LastName)
		assert.Equal(T, user.Email, retUser.Email)
	})

	Given(`^two aggregates$`, func() {
		user2,_ = sample.NewUser("first2","last2", "email2")
		user2.UpdateFirstName("new first")
		assert.Equal(T, "new first", user2.FirstName)
	})

	And(`^an event store containing all events for both aggregates$`, func() {
		err := user2.Store(eventStore)
		assert.Nil(T,err)
		assert.Equal(T, 0, len(user2.Events))
	})

	When(`^the events for an aggregate are retrieved$`, func() {
		eventSets, err := eventStore.RetrieveEvents(user2.ID)
		assert.Nil(T,err)
		if assert.Equal(T, 1, len(eventSets), "Expected one event set to be retrieved") {
			assert.Equal(T, 2, len(eventSets[0]))
		}

	})

	Then(`^only the events associated with the specific aggregate are retrieved$`, func() {
		eventSets, err := eventStore.RetrieveEvents(user2.ID)
		assert.Nil(T,err)
		retUser := sample.NewUserFromHistory(eventSets[0])
		assert.NotNil(T, retUser)
		assert.Equal(T, user2.FirstName, retUser.FirstName)
		assert.Equal(T, user2.LastName, retUser.LastName)
		assert.Equal(T, user2.Email, retUser.Email)
	})


}
