package es2


type EventStore interface {
	StoreEvents(*Aggregate) error
	RetrieveEvents(aggID string) ([]Event,error)
}


type EventSourced interface  {
	Store(EventStore)
	Apply(Event)
	Route(Event)
}