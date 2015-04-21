package main 

import (
	"net/http"
	"net/http/httptest"
)

func handleCall(w http.ResponseWriter, r *http.Request) {
	println("handleCall called...")
	w.Write([]byte("handleCall wrote this stuff\n"))
}

func wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()
		println("wrap...")
		h.ServeHTTP(rec, r)
		w.Write(rec.Body.Bytes())
		w.Write([]byte("wrap wrote this\n"))	
 	})	
}

func main() {
	println("yeah")
	handleCallHandler := http.HandlerFunc(handleCall)
	wrapped := wrap(handleCallHandler)
	http.ListenAndServe(":8080", wrapped)
}



