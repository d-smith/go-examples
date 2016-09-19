package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func makeClient() *http.Client {
	certs := x509.NewCertPool()

	pemData, err := ioutil.ReadFile("../cert.pem")
	if err != nil {
		fatal(err)
	}
	certs.AppendCertsFromPEM(pemData)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: certs},
	}

	return &http.Client{Transport: tr}
}

func main() {

	client := makeClient()

	req, err := http.NewRequest("GET", "https://localhost:10443", nil)
	fatal(err)

	resp, err := client.Do(req)
	fatal(err)

	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	fatal(err)

	log.Println(string(all))
}
