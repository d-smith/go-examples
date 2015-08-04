package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/samalba/dockerclient"
	"io/ioutil"
	"log"
	"os"
)

func readDockerEnv() (string, string) {

	dockerHome := os.Getenv("DOCKER_HOST")
	if dockerHome == "" {
		log.Fatal("DOCKER_HOST environment variable not set.")
	}

	dockerCertPath := os.Getenv("DOCKER_CERT_PATH")
	if dockerCertPath == "" {
		log.Fatal("DOCKER_CERT_PATH environment variable not set.")
	}

	return dockerHome, dockerCertPath
}

func buildDockerTLSConfig(dockerCertPath string) *tls.Config {

	caFile := fmt.Sprintf("%s/ca.pem", dockerCertPath)
	certFile := fmt.Sprintf("%s/cert.pem", dockerCertPath)
	keyFile := fmt.Sprintf("%s/key.pem", dockerCertPath)

	tlsConfig := &tls.Config{}

	cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
	pemCerts, _ := ioutil.ReadFile(caFile)

	tlsConfig.RootCAs = x509.NewCertPool()
	tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	tlsConfig.Certificates = []tls.Certificate{cert}

	tlsConfig.RootCAs.AppendCertsFromPEM(pemCerts)

	return tlsConfig
}

func main() {
	//Grab the environment
	dockerHost, dockerCertPath := readDockerEnv()

	// Init the client
	log.Println("Create docker client")
	docker, _ := dockerclient.NewDockerClient(dockerHost, buildDockerTLSConfig(dockerCertPath))

	containers, err := docker.ListContainers(true, true, "")
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range containers {
		fmt.Printf("%v\n", c)
		ci,_ := docker.InspectContainer(c.Id)
		fmt.Printf("\t%v\n", ci)
		links := ci.HostConfig.Links
		fmt.Printf("\t--links: %v\n", links)
		fmt.Printf("\t--labels: %v\n", c.Labels)
	}

}
