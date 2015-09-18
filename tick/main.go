package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//
// Note: beware resource leaks when using tick. See the time package docs for more info.
//

var wg sync.WaitGroup

func intercomReceiver(intercom chan string, done chan string) {
	defer wg.Done()

	for {
		select {
		case msg, ok := <-intercom:
			if !ok {
				fmt.Println("channel closed - shutting down")
				return
			}

			fmt.Println("received ", msg)
		case <-time.After(5 * time.Second):
			done <- "Laters"
			return
		}
	}
}

func intercomSender(intercom chan string, done chan string, maxdelay int) {

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGen := rand.New(randomSource)

	defer wg.Done()
	for i := 0; i < 50; i++ {
		select {
		case intercom <- fmt.Sprintf("message %d", i):
			delay := randomGen.Intn(maxdelay + 1)
			time.Sleep(time.Duration(delay) * time.Second)
		case quitmsg := <-done:
			fmt.Println("receiver quit: ", quitmsg)
			return
		}
	}

	close(intercom)
}

func main() {
	var maxdelay = flag.Int("maxdelay", -1, "maximum delay between message sends")
	flag.Parse()
	if *maxdelay == -1 {
		fmt.Println("Must specify a -maxdelay argument")
		return
	}

	//Create a done channel to allow the receiver to cancel any further messaging
	done := make(chan string)

	//Create the channel for sending things to the receiver
	intercom := make(chan string)
	wg.Add(2)

	go intercomReceiver(intercom, done)
	go intercomSender(intercom, done, *maxdelay)

	wg.Wait()
}
