package main

import (
	"github.com/hydrogen18/test-tls"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
)


func main() {
	if len(os.Args) < 4 {
		fmt.Println("Need three args: client key file, client cert file, ca cert file")
		os.Exit(1)
	}

	log.Println("get tls config")
	config := common.MustGetTlsConfiguration()

	log.Println("create client with tls transport config")
	tr := &http.Transport{
		TLSClientConfig:config,
	}

	client := http.Client{Transport: tr}

	log.Println("create request")
	req, err := http.NewRequest("GET", "https://localhost:5000/svc", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Get /")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	log.Println("Read response")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("status: %d", resp.StatusCode)
	log.Printf("response: %s", string(body))

}
