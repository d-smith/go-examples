package eventstore

import (
	"errors"
	"github.com/d-smith/go-examples/es2"
	"sync"
)

//TODO - package error pattern - const or var??

type subscriberStorage struct {
	subId    es2.SubscriptionID
	callback es2.EventPublishedCallback
}

type EventStorage struct {
	events         []es2.Event
	currentVersion int
}

type InMemoryEventStore struct {
	sync.RWMutex
	storage     map[string]EventStorage
	subscribers []subscriberStorage
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		storage: make(map[string]EventStorage),
	}
}

func (im *InMemoryEventStore) publishEvent(event es2.Event) {
	for _, sub := range im.subscribers {
		sub.callback(event)
	}
}

func (im *InMemoryEventStore) StoreEvents(agg *es2.Aggregate) error {
	im.Lock()
	defer im.Unlock()

	//Do we have events for this aggregate?
	aggStorage, ok := im.storage[agg.ID]
	if !ok {
		aggStorage = EventStorage{}
	}

	//Has someone update the aggregate before the current caller?
	if !(aggStorage.currentVersion < agg.Version) {
		return errors.New("Concurrency exception")
	}

	//Set the new version, and append the events
	aggStorage.currentVersion = agg.Version
	for _, e := range agg.Events {
		aggStorage.events = append(aggStorage.events, e)
		im.publishEvent(e)
	}

	im.storage[agg.ID] = aggStorage

	return nil
}

func (im *InMemoryEventStore) RetrieveEvents(aggId string) ([]es2.Event, error) {
	im.RLock()
	defer im.RUnlock()

	eventStorage, ok := im.storage[aggId]
	if !ok {
		return nil, errors.New("No events stored for aggregate")
	}

	return eventStorage.events, nil
}

func (im *InMemoryEventStore) SubscribeEvents(callback es2.EventPublishedCallback) es2.SubscriptionID {
	im.Lock()
	defer im.Unlock()
	subId := es2.SubscriptionID(es2.GenerateID())
	im.subscribers = append(im.subscribers, subscriberStorage{subId: subId, callback: callback})
	return subId

}

func (im *InMemoryEventStore) Unsubscribe(subId es2.SubscriptionID) {
	im.Lock()
	remainingSubs := make([]subscriberStorage, 0, len(im.subscribers)-1)
	for _, sub := range im.subscribers {
		if sub.subId != subId {
			remainingSubs = append(remainingSubs, sub)
		}
	}
	im.subscribers = remainingSubs
	im.Unlock()
}
