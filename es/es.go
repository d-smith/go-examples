package es

import (
	"errors"
	"github.com/nu7hatch/gouuid"
	"log"
	"sync"
)

//Event contains the aggregate id the event is associated with and the
//event payload
type Event struct {
	AggregateId string
	Payload     interface{}
}

//EventStore defines the interface an Event Store must provide.
type EventStore interface {
	GetEvents(aggregateId string) ([]Event, error)
	StoreEvents(aggregateId string, events []Event) error
}

type InMemEventStore struct {
	sync.RWMutex
	AllEvents map[string][]Event
}

func NewInMemEventStore(eventStream EventStream) *InMemEventStore {
	es := new(InMemEventStore)
	es.AllEvents = make(map[string][]Event)
	go func() {
		for {
			log.Println("reading from event stream...")
			event := eventStream.Get()
			log.Println("read event", event)
			es.StoreEvents([]Event{event})
		}
	}()
	return es
}

func (es *InMemEventStore) GetEvents(aggregateId string) ([]Event, error) {
	es.RLock()
	defer es.RUnlock()
	events, ok := es.AllEvents[aggregateId]
	if !ok {
		return nil, errors.New("No events for aggregateId " + aggregateId)
	}

	return events, nil
}

func (es *InMemEventStore) StoreEvents(events []Event) error {
	es.Lock()
	defer es.Unlock()

	for _, event := range events {
		//Get the current set of events for the aggregate id
		allEvents := es.AllEvents[event.AggregateId]
		allEvents = append(allEvents, event)
		es.AllEvents[event.AggregateId] = allEvents
	}

	return nil
}

type EventSourced struct {
	AggregateId string
}

type EventRecorder interface {
	Record(event Event)
	Flush(es EventStream)
}

type LocalEventRecorder struct {
	aggregateId string
	events      []Event
}

func NewLocalEventRecorder(aggregateId string) *LocalEventRecorder {
	return &LocalEventRecorder{
		aggregateId: aggregateId,
		events:      make([]Event, 0),
	}
}

func (er *LocalEventRecorder) Record(event Event) {
	log.Println("record event", event)
	er.events = append(er.events, event)
}

func (er *LocalEventRecorder) Flush(es EventStream) {
	for _, event := range er.events {
		log.Println("flush event", event)
		es.Put(event)
	}
}

func GenerateID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

type EventStream interface {
	Put(event Event)
	Get() Event
}

type InMemEventStream struct {
	stream chan Event
}

func NewInMemEventStream() *InMemEventStream {
	return &InMemEventStream{
		stream: make(chan Event, 50),
	}
}

func (es *InMemEventStream) Get() Event {
	return <-es.stream
}

func (es *InMemEventStream) Put(event Event) {
	es.stream <- event
}
