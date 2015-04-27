package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/stuff/", http.StripPrefix("/stuff/", fs))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
