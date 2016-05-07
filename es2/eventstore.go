package es2

type EventSet []Event

type EventStore interface {
	StoreEvents(*Aggregate) error
	RetrieveEvents(aggID string) ([]EventSet,error)
}


type EventSourced interface  {
	Store(EventStore)
	Apply(Event)
	Route(Event)
}