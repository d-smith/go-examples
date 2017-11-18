package main

import (
	"fmt"
	"net/http"
	"os"
)

func buildPackHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hey yo from cf"))
}


func main() {
	http.HandleFunc("/", buildPackHandler)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

