package main

import (
	"net/http"
)

func handler1(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("handler 1"))
}

func handler2(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("handler 2"))
}

func handler3(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("handler 3"))
}

func main() {
	http.HandleFunc("/foo", handler1)
	http.HandleFunc("/foo", handler2)
	http.HandleFunc("/foo", handler3)

	http.ListenAndServe(":8080", nil)
}
