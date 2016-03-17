package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	log.Println("Handler invoked")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello world\n"))
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
