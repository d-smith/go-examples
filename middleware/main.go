package main 

import (
	"net/http"
)

func handleCall(w http.ResponseWriter, r *http.Request) {
	println("handleCall called...")
	w.Write([]byte("handleCall wrote this stuff\n"))
}

func wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("wrap...")
		h.ServeHTTP(w, r)
		w.Write([]byte("wrap wrote this\n"))	
 	})	
}

func main() {
	println("yeah")
	handleCallHandler := http.HandlerFunc(handleCall)
	wrapped := wrap(handleCallHandler)
	http.ListenAndServe(":8080", wrapped)
}



