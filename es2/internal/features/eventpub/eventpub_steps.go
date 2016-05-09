package eventpub

import (
	"github.com/d-smith/go-examples/es2"
	"github.com/d-smith/go-examples/es2/eventstore"
	"github.com/d-smith/go-examples/es2/sample"
	. "github.com/lsegal/gucumber"
	"github.com/stretchr/testify/assert"
)

func init() {

	var user *sample.User
	var events []es2.Event
	var eventStore es2.EventStore = eventstore.NewInMemoryEventStore()
	var subId es2.SubscriptionID

	var callback = func(event es2.Event) {
		events = append(events, event)
	}

	When(`^I create and modify an instance of the aggregate$`, func() {
		user, _ = sample.NewUser("first", "last", "email")
		user.UpdateFirstName("updated")
		subId = eventStore.SubscribeEvents(callback)
		user.Store(eventStore)
	})

	Then(`^all the events are published$`, func() {
		assert.Equal(T, 2, len(events))
	})

	Then(`^no events are published$`, func() {
		eventHistory, _ := eventStore.RetrieveEvents(user.ID)
		sample.NewUserFromHistory(eventHistory)
		assert.Equal(T, 2, len(events))
	})

	Given(`^an event store with a registered subscriber$`, func() {
	})

	When(`^the subscriber unsubscribes$`, func() {
		eventStore.Unsubscribe(subId)
	})

	Then(`^previously subscribed callback is not invoked when events are published$`, func() {
		user, _ = sample.NewUser("first", "last", "email")
		user.Store(eventStore)
		assert.Equal(T, 2, len(events))
	})

}
