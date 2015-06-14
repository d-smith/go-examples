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

func main() {
	a := actor.NewActor(myReceiver)
	go a.Run()
	var i int
	for {
		a.Send(i, nil)
		i++
	}
}
