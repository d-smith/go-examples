package main

import (
	"net/http"
	"strings"
	"log"
)

var callCount = 0

func echoHandler(rw http.ResponseWriter, req *http.Request) {
	uriParts := strings.Split(req.RequestURI, "/")
	rw.Write([]byte(uriParts[len(uriParts)-1]))
	rw.Write([]byte("\n"))
}

func healthHandler(rw http.ResponseWriter, req *http.Request) {
	callCount++
	log.Println("health check called")

	if callCount > 10 {
		http.Error(rw, "I am broken", http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo/", echoHandler)
	mux.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":4000", mux)
}