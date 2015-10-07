package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var finishedChannel = make(chan bool)
var stopped = false

func doStuff() {
	for stopped == false {
		fmt.Println("Doing some stuff...")
		time.Sleep(2 * time.Second)
	}
	fmt.Println("shutdown complete")
	finishedChannel <- true
}

func main() {
	go doStuff()

	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		for _ = range signalChannel {
			fmt.Println("Received interrupt - shutting down")
			stopped = true
		}
	}()
	<-finishedChannel

}
