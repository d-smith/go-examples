package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func makeClient() *http.Client {
	certs := x509.NewCertPool()

	pemData, err := ioutil.ReadFile("../cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	certs.AppendCertsFromPEM(pemData)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: certs},
	}

	return &http.Client{Transport: tr}
}

func main() {
	client := makeClient()
	host, _ := os.Hostname()
	for {
		req, err := http.NewRequest("GET", "https://"+host+":4000/feed", nil)
		if err != nil {
			log.Fatal(err.Error())
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("proto version", resp.ProtoMajor)

		defer resp.Body.Close()

		rb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println(string(rb))

		time.Sleep(5 * time.Second)
	}
}
