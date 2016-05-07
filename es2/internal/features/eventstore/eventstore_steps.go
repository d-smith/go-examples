package eventstore

import (
	. "github.com/lsegal/gucumber"
	"github.com/d-smith/go-examples/es2/sample"
	"github.com/d-smith/go-examples/es2/eventstore"
	"github.com/d-smith/go-examples/es2"
	"github.com/stretchr/testify/assert"
)

func init() {
	var user *sample.User
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
		T.Skip() // pending
	})

	And(`^an event store containing all events for both aggregates$`, func() {
		T.Skip() // pending
	})

	When(`^the events for an aggregate are retrieved$`, func() {
		T.Skip() // pending
	})

	Then(`^only the events assocaited with the specific aggregate are retrieved$`, func() {
		T.Skip() // pending
	})

	And(`^the events are ordered$`, func() {
		T.Skip() // pending
	})

}
