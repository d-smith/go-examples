package main

import (
	"net/http"
	"strings"
)

func echoHandler(rw http.ResponseWriter, req *http.Request) {
	uriParts := strings.Split(req.RequestURI, "/")
	rw.Write([]byte(uriParts[len(uriParts)-1]))
	rw.Write([]byte("\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo/", echoHandler)
	http.ListenAndServe(":4000", mux)
}
