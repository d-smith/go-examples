package sample

import (
	"github.com/d-smith/go-examples/es2"
	"log"
)

//To be event sourced...
// The Aggregate type must be embedded
// A constructor for a brand new aggregate is available
// Constructing an aggregate produces an event

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
			Payload: UserCreated{
				AggregateId: user.Aggregate.ID,
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
	switch event.Payload.(type) {
	case UserCreated:
		u.handleUserCreated(event.Payload.(UserCreated))
	case UserFirstNameUpdated:
		u.handleUserFirstNameUpdate(event.Payload.(UserFirstNameUpdated))
	default:
		log.Println("WARN: unknown event routed to User aggregate")
	}
}

func (u *User) Apply(event es2.Event) {
	u.Route(event)
	u.Aggregate.Events = append(u.Aggregate.Events, event)
}
