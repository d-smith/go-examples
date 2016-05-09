package sample

import (
	"github.com/d-smith/go-examples/es2"
	"log"
)

//To be event sourced...
// The Aggregate type must be embedded
// A constructor for a brand new aggregate is available
// Constructing an aggregate produces an event
// Mutations occur via commands, which emit events handled by event handler, with events routed to handlers
// Events are recorded in event history
// An apply method routes an event to the event handler, and records the event
// When applying event history, only the route method is used - side effects occur in the command handlers

type User struct {
	*es2.Aggregate
	FirstName string
	LastName  string
	Email     string
}

func NewUser(first, last, email string) (*User, error) {
	//Do validation... return an error if there's a problem
	var user = new(User)
	user.Aggregate = es2.NewAggregate()

	user.Apply(
		es2.Event{
			Source: user.ID,
			Payload: UserCreated{
				AggregateId: user.ID,
				FirstName:   first,
				LastName:    last,
				Email:       email,
			},
		})

	return user, nil
}

func NewUserFromHistory(events []es2.Event) *User {
	user := new(User)
	user.Aggregate = es2.NewAggregate()

	for _, e := range events {
		log.Println("apply event", e)
		user.Route(e)
	}

	return user
}

type UserCreated struct {
	AggregateId string
	FirstName   string
	LastName    string
	Email       string
}

type UserFirstNameUpdated struct {
	OldFirst string
	NewFirst string
}

type UserLastNameUpdated struct {
	OldLast string
	NewLast string
}

func (u *User) UpdateFirstName(first string) {
	u.Apply(
		es2.Event{
			Source: u.ID,
			Payload: UserFirstNameUpdated{
				OldFirst: u.FirstName,
				NewFirst:   first,
			},
		})
}

func (u *User) handleUserCreated(event UserCreated) {
	u.Aggregate.ID = event.AggregateId
	u.FirstName = event.FirstName
	u.LastName = event.LastName
	u.Email = event.Email
}

func (u *User) handleUserFirstNameUpdate(event UserFirstNameUpdated) {
	u.FirstName = event.NewFirst
}

func (u *User) Route(event es2.Event) {
	u.Version += 1
	event.Version = u.Version
	switch event.Payload.(type) {
	case UserCreated:
		u.handleUserCreated(event.Payload.(UserCreated))
	case UserFirstNameUpdated:
		u.handleUserFirstNameUpdate(event.Payload.(UserFirstNameUpdated))
	default:
		log.Println("WARN: unknown event routed to User aggregate")
		u.Version -= 1
	}
}

func (u *User) Apply(event es2.Event) {
	u.Route(event)
	u.Aggregate.Events = append(u.Aggregate.Events, event)
}

func (u *User) Store(eventStore es2.EventStore) error {
	err := eventStore.StoreEvents(u.Aggregate)
	if err != nil {
		return err
	}

	u.Events = make([]es2.Event,0)

	return nil
}
