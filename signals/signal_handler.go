package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func doStuff() {
	for {
		fmt.Println("Doing some stuff...")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	go doStuff()
	
	signalChannel := make(chan os.Signal, 1)
	finishedChannel := make(chan bool)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		for _ = range signalChannel {
			fmt.Println("Received interrupt - shutting down")
			finishedChannel <- true
		}
	}()
	<- finishedChannel
	
}

