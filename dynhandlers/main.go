package main

import (
	"net/http"
)

var mux *http.ServeMux
var urisMap map[string]string

func doPut(rw http.ResponseWriter, req *http.Request) {
	uri := req.FormValue("uri")
	if uri == "" {
		rw.Write([]byte("I couldn't parse your uri\n"))
		return
	}

	registered := urisMap[uri]
	if registered != "" {
		rw.Write([]byte("URI already registered\n"))
		return
	}

	urisMap[uri] = uri
	mux.HandleFunc("/"+uri, func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("This is the " + uri + " handler\n"))
	})
}

func entryHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		rw.Write([]byte("Try put with single uri query param\n"))
		break
	case http.MethodPut:
		doPut(rw, req)
		break
	default:
		http.Error(rw, "I only support GET and PUT\n", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux = http.NewServeMux()
	urisMap = make(map[string]string)

	mux.HandleFunc("/register", entryHandler)

	http.ListenAndServe(":4000", mux)
}
