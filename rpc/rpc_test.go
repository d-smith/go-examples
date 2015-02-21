package main

import (
	"fmt"
	"time"
	"testing"
)

func TestPing(t *testing.T) {
	
	pc,err := NewPingClient(":9876", 5*time.Second)
	if err != nil {
		msg := "Skipping test - no connction to RPC server"
		fmt.Println(msg)
		t.Skip(msg)
	}
	
	out, err := pc.Ping("this is a test")
	if err != nil {
		t.Error("Error calling ping",err.Error())
	}
	
	if out != "THIS IS A TEST" {
		t.Error("Did not get all caps back: ", out)
	}
}
