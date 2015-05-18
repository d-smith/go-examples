package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're all boring; I'm leaving.")
}

// Returns receive-only channel of strings.
func boring(msg string) <-chan Message {
	c := make(chan Message)

	// Shared between all messages.
	waitForIt := make(chan bool)

	// We launch the goroutine from inside the function.
	go func() {
		for i := 0; ; i++ {

			c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt

		}
	}()

	// Return the channel to the caller.
	return c
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		// New instance of 'input' for each loop.
		input := inputs[i]
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}
