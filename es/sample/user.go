package main

import (
	"errors"
	"github.com/d-smith/go-examples/es"
)

var myEventStream = es.NewInMemEventStream()
var myEventStore = es.NewInMemEventStore(myEventStream)

//Domain object

type User struct {
	es.EventSourced
	es.EventRecorder
	FirstName string
	LastName  string
	Email     string
}

//Commands

func (u *User) UpdateFirstName(first string) error {
	if first == "" {
		return errors.New("First name must not be empty")
	}

	return u.Apply(
		es.Event{
			AggregateId: u.AggregateId,
			Payload: UserFirstNameUpdated{
				OldFirst: u.FirstName,
				NewFirst: first,
			},
		})
}

func (u *User) Apply(event es.Event) error {
	err := u.Route(event)
	if err == nil {
		u.Record(event)
	}

	return err
}

func (u *User) UpdateLastName(last string) {

}

//Events

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

// Handlers

func (u *User) handleUserCreated(event UserCreated) error {
	var err error
	u.FirstName = event.FirstName
	u.LastName = event.LastName
	u.Email = event.Email
	u.AggregateId = event.AggregateId
	return err
}

func (u *User) handleUserFirstNameUpdated(event UserFirstNameUpdated) error {
	u.FirstName = event.NewFirst
	return nil
}

func (u *User) handleUserLastNameUpdate() {

}

// Router

func (u *User) Route(event es.Event) error {
	switch event.Payload.(type) {
	case UserCreated:
		return u.handleUserCreated(event.Payload.(UserCreated))
	case UserFirstNameUpdated:
		return u.handleUserFirstNameUpdated(event.Payload.(UserFirstNameUpdated))
	default:
		return errors.New("event type not recognized	")
	}
}

// Constructor - new
func NewUser(first, last, email string) (*User, error) {
	//Do validation... then
	user := new(User)
	aggId, err := es.GenerateID()
	if err != nil {
		return nil, err
	}

	user.EventRecorder = es.NewLocalEventRecorder(aggId)

	err = user.Apply(
		es.Event{
			AggregateId: aggId,
			Payload: UserCreated{
				AggregateId: aggId,
				FirstName:   first,
				LastName:    last,
				Email:       email,
			},
		})

	return user, err
}

//Constructor - from events
func NewUserFromHistory(events []es.Event) (*User, error) {
	user := new(User)

	for _, e := range events {
		err := user.Route(e)
		if err != nil {
			return nil, err
		}
	}

	user.EventRecorder = es.NewLocalEventRecorder(user.AggregateId)

	return user, nil
}
