package main

import (
	"log"
	"net/http"
	"net/http/httptest"
)

var responseJson string = `
{
  "name": "GOOGL",
  "last": 1002.20,
  "time": "12:34",
  "date": "10/31/2014"
}
`

func handleQuoteCalls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Write([]byte(responseJson))
}

func corsWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, r)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(rec.Body.Bytes())
	})
}

func main() {
	//Original static content
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/stuff/", http.StripPrefix("/stuff/", fs))

	//Add some dynamc content
	http.HandleFunc("/quote/", handleQuoteCalls)

	//Add the swagger spec
	ss := http.FileServer(http.Dir("dist"))
	http.Handle("/apispec/", http.StripPrefix("/apispec/", corsWrapper(ss)))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
