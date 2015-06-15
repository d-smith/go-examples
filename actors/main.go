package main

import (
	"fmt"
	"github.com/d-smith/go-examples/actors/actor"
)

func myReceiver(rc chan actor.Message) {
	msg := <-rc
	switch msg.Content.(type) {
	case int:
		fmt.Printf("received an integer: %d\n", msg.Content.(int))
	default:
		fmt.Printf("received value %v of type %T\n", rc, rc)
	}
}

type MyActor struct {
	*actor.Actor
	receiveFn func(chan actor.Message)
}

func main() {
	myActor := &MyActor{
		Actor: &actor.Actor{
			Mailbox:   make(chan actor.Message, 100),
			ReceiveFn: myReceiver,
		},
	}

	go myActor.Run()
	var i int
	for {
		myActor.Send(i, nil)
		i++
	}
}
