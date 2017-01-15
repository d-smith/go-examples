package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// Load client cert
	cert, err := tls.LoadX509KeyPair("okguy.crt", "okguy.key")
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile("../server/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	resp, err := client.Get("https://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	} 
	
	fmt.Printf("%s\n", string(contents))

	peerCerts := resp.TLS.PeerCertificates
	for _, pc := range peerCerts {
		log.Printf("organization: %s", pc.Subject.Organization)
		log.Printf("common name: %s", pc.Subject.CommonName)
	}
}
