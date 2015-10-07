package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Example adapted from the go documentation. Used to verify use the http_proxy
//environment variable is picked up when making calls in a proxied environment.
func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	resource := os.Args[1]

	res, err := http.Get(resource)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
