package main

import (
	"errors"
	"github.com/nu7hatch/gouuid"
	"sync"
)

type EventStore interface {
	GetEvents(aggregateId string) ([]interface{}, error)
	StoreEvents(aggregateId string, events []interface{}) error
}

type InMemEventStore struct {
	sync.RWMutex
	AllEvents map[string][]interface{}
}

func NewInMemEventStore() *InMemEventStore {
	es := new(InMemEventStore)
	es.AllEvents = make(map[string][]interface{})
	return es
}

func (es *InMemEventStore) GetEvents(aggregateId string) ([]interface{}, error) {
	es.RLock()
	defer es.RUnlock()
	events, ok := es.AllEvents[aggregateId]
	if !ok {
		return nil, errors.New("No events for aggregateId " + aggregateId)
	}

	return events, nil
}

var myEventStore = NewInMemEventStore()

func (es *InMemEventStore) StoreEvents(aggregateId string, events []interface{}) error {
	es.Lock()
	defer es.Unlock()

	//Get the current set of events
	allEvents := es.AllEvents[aggregateId]

	//Append the new set of events
	for _, event := range events {
		allEvents = append(allEvents, event)
	}

	es.AllEvents[aggregateId] = allEvents

	return nil
}

type EventSourced struct {
	AggregateId string
	Events      []interface{}
}

//Domain object

type User struct {
	EventSourced
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
	err := u.Route(event)
	if err == nil {
		myEventStore.StoreEvents(u.AggregateId, []interface{}{event})
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
	var err error
	u.FirstName = event.FirstName
	u.LastName = event.LastName
	u.Email = event.Email
	u.AggregateId, err = GenerateID()
	return err
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
func NewUserFromHistory(events []interface{}) (*User, error) {
	user := new(User)
	for _, e := range events {
		err := user.Route(e)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

type EventRecorder struct {
	events []interface{}
}

func (er *EventRecorder) Record(event interface{}) {
	er.events = append(er.events, event)
}

func (er *EventRecorder) GetEvents() []interface{} {
	return er.events
}

func GenerateID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
