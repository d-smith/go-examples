package es2

type EventPublishedCallback func(event Event)

type SubscriptionID string

type EventStore interface {
	StoreEvents(*Aggregate) error
	RetrieveEvents(aggID string) ([]Event, error)
	SubscribeEvents(callback EventPublishedCallback) SubscriptionID
	Unsubscribe(sub SubscriptionID)
}

type EventSourced interface {
	Store(EventStore)
	Apply(Event)
	Route(Event)
}
