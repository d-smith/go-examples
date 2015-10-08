package main

import (
	"net/http"
	"log"
)

type handler func(w http.ResponseWriter, r *http.Request)

func (h handler) decorate(hf handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		hf(w,r)
		h(w,r)
	}
}

var invocationLogger = func (w http.ResponseWriter, r *http.Request) {
	log.Println("invoked for", r.URL)
}

var handleFoo handler = func(w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("got it"))
}

func main() {
	decorated := handleFoo.decorate(invocationLogger)
	http.Handle("/foo", http.HandlerFunc(decorated))
	http.ListenAndServe(":8080", nil)
}
