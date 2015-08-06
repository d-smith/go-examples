package main

import (
	"net"
	"log"
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var wg sync.WaitGroup

var services = []string {
	"createWorkItem",
	"updateWorkItem",
	"retieveWorkItemRule",
	"findWorkItem",
	"lockWorkItem",
	"retrieveWorkItem",
	"attachDocument",
	"splitItem",
}

var apps = []string {
"app01", "app02","app03","app04","app05", "app06",
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleConnection(conn net.Conn) {
	for  {
		println("write to connection")
		logLine := fmt.Sprintf("%s %s", apps[rand.Intn(len(apps))], services[rand.Intn(len(services))])
		fmt.Fprintln(conn, logLine)
		time.Sleep(200 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	println("listen for connection")
	ln, err := net.Listen("tcp", ":9999")
	fatal(err)

	println("accept connection")
	conn, err := ln.Accept()
	fatal(err)
	println("handle connection")
	go handleConnection(conn)

	wg.Wait()

}
