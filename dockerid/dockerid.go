package main

import (
	"log"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Hostname :%s", name)
	log.Printf("PID: %d", os.Getpid())
}
