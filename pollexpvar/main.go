package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	for {
		dumpExpvar()
		time.Sleep(10 * time.Second)
	}

}

func dumpExpvar() {
	resp, err := http.Get("http://localhost:8080/debug/vars")
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	switch err {
	case nil:
		log.Println(string(data))

	default:
		log.Println(err.Error())

	}

}
