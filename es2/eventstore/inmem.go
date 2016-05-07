package eventstore

import (
	"github.com/d-smith/go-examples/es2"
	"sync"
	"errors"
)

//TODO - package error pattern - const or var??

type InMemoryEventStore struct {
	sync.RWMutex
	storage map[string][]es2.EventSet
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		storage:make(map[string][]es2.EventSet),
	}
}

func (im *InMemoryEventStore) StoreEvents(agg *es2.Aggregate) error {
	im.Lock()
	defer im.Unlock()
	aggEvents := im.storage[agg.ID]
	if aggEvents == nil {
		aggEvents = make([]es2.EventSet,0)
	}

	aggEvents = append(aggEvents, agg.Events)
	im.storage[agg.ID] = aggEvents

	return nil
}

func (im *InMemoryEventStore) RetrieveEvents(aggId string) ([]es2.EventSet,error) {
	im.RLock()
	defer im.RUnlock()

	eventSet, ok := im.storage[aggId]
	if !ok {
		return nil, errors.New("No events stored for aggregate")
	}

	return eventSet, nil
}
