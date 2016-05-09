package eventstore

import (
	"errors"
	"github.com/d-smith/go-examples/es2"
	"sync"
)

//TODO - package error pattern - const or var??

type EventStorage struct {
	events         []es2.Event
	currentVersion int
}

type InMemoryEventStore struct {
	sync.RWMutex
	storage map[string]EventStorage
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		storage: make(map[string]EventStorage),
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
