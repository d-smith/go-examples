package main

import (
	"errors"
)

/*
Implementation map

Define:
	user (domain obj)
	commands
	events
	handlers - route and record
	routing (event to handler)
	recording (of events)
	constructor - new domain object
	constructor - event history

	Next - need a way to record history, and a way to load history which means we'll need
	an aggregate id, and an event store.
*/

//Domain object

type User struct {
	FirstName string
	LastName  string
	Email     string
}

//Commands

func (u *User) UpdateFirstName(first string) error {
	if first == "" {
		return errors.New("First name must not be empty")
	}

	return u.Apply(UserFirstNameUpdated{
		OldFirst: u.FirstName,
		NewFirst: first,
	})
}

func (u *User) Apply(event interface{}) error {
	err :=  u.Route(event)
	if err == nil {
		eventStore.Record(event)
	}

	return err
}

func (u *User) UpdateLastName(last string) {

}

//Events

type UserCreated struct {
	FirstName string
	LastName  string
	Email     string
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
	u.FirstName = event.FirstName
	u.LastName = event.LastName
	u.Email = event.Email
	return nil
}

func (u *User) handleUserFirstNameUpdated(event UserFirstNameUpdated) error {
	u.FirstName = event.NewFirst
	return nil
}

func (u *User) handleUserLastNameUpdate() {

}

// Router

func (u *User) Route(event interface{}) error {
	switch event.(type) {
	case UserCreated:
		return u.handleUserCreated(event.(UserCreated))
	case UserFirstNameUpdated:
		return u.handleUserFirstNameUpdated(event.(UserFirstNameUpdated))
	default:
		return errors.New("event type not recognized	")
	}
}

// Constructor - new
func NewUser(first, last, email string) (*User, error) {
	//Do validation... then
	user := new(User)
	err := user.Apply(
		UserCreated{
		FirstName: first,
		LastName:  last,
		Email:     email,
	})


	return user, err
}

//Constructor - from events
func NewUserFromHistory(events []interface{}) (*User,error) {
	user := new(User)
	for _,e := range events {
		err := user.Route(e)
		if err != nil {
			return nil,err
		}
	}

	return user, nil
}

type EventRecorder struct {
	events []interface{}
}

var eventStore = new(EventRecorder)

func (er *EventRecorder) Record(event interface{}) {
	er.events = append(er.events, event)
}

func (er *EventRecorder) GetEvents() []interface{} {
	return er.events
}
