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

type UserCreated struct {
	AggregateId string
	FirstName   string
	LastName    string
	Email       string
}

func (u *User) handleUserCreated(event UserCreated) {
	u.Aggregate.ID = event.AggregateId
	u.FirstName = event.FirstName
	u.LastName = event.LastName
	u.Email = event.Email
}

func (u *User) Route(event es2.Event) {
	switch event.Payload.(type) {
	case UserCreated:
		u.handleUserCreated(event.Payload.(UserCreated))
	default:
		log.Println("WARN: unknown event routed to User aggregate")
	}
}

func (u *User) Apply(event es2.Event) {
	u.Route(event)
	u.Aggregate.Events = append(u.Aggregate.Events, event)
}
