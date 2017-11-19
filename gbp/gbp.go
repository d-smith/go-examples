package main

import (
	"fmt"
	"net/http"
	"os"
)

func buildPackHandler(res http.ResponseWriter, req *http.Request) {

	for _, e := range os.Environ() {
		res.Write([]byte(fmt.Sprintf("%s\n",e)))
	}
}


func main() {
	http.HandleFunc("/", buildPackHandler)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

