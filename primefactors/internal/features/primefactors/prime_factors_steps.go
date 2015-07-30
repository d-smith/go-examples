package primefactors

import (
	"log"
	. "github.com/lsegal/gucumber"
	"github.com/samalba/dockerclient"
	"io/ioutil"
	"crypto/tls"
	"crypto/x509"
	"os"
	"fmt"
)

func readEnv() (string,string) {

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

func buildTLSConfig(dockerCertPath string) *tls.Config {


	caFile := fmt.Sprintf("%s/ca.pem", dockerCertPath)
	certFile := fmt.Sprintf("%s/cert.pem", dockerCertPath)
	keyFile := fmt.Sprintf("%s/key.pem", dockerCertPath)

	tlsConfig := &tls.Config{}

	cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
	pemCerts, _ := ioutil.ReadFile(caFile)

	tlsConfig.RootCAs       = x509.NewCertPool()
	tlsConfig.ClientAuth    = tls.RequireAndVerifyClientCert
	tlsConfig.Certificates  = []tls.Certificate{cert}

	tlsConfig.RootCAs.AppendCertsFromPEM(pemCerts)

	return tlsConfig
}

func getAcceptanceTestContainerInfo(docker *dockerclient.DockerClient, containerType string) *dockerclient.ContainerInfo {

	// Get only running containers
	containers, err := docker.ListContainers(false, false, "")
	if err != nil {
		log.Fatal(err)
	}

	//Loop through them until we find a match
	for _, c := range containers {
		xtContainerType, ok := c.Labels["xt-container-type"]
		if ok && xtContainerType == containerType {
			//Grab the information for the container
			info, err := docker.InspectContainer(c.Id)
			if err != nil {
				log.Fatal(err)
			}

			return info
		}
	}

	return nil
}

func init() {

	var containerId string
	var docker *dockerclient.DockerClient

	Before("@primefactors", func() {
		//Grab the environment
		dockerHost, dockerCertPath := readEnv()

		// Init the client
		docker, _ = dockerclient.NewDockerClient(dockerHost, buildTLSConfig(dockerCertPath))

		info := getAcceptanceTestContainerInfo(docker, "atest")
		if info != nil {
			log.Println("Container found - state is: ", info.State.StateString())
			log.Println("Test is ready to run")
			return
		}

		//No running container found - start one. This assumes the image we want to run is available
		//to the docker runtime.
		labels := make(map[string]string)
		labels["xt-container-type"] = "atest"
		containerConfig := dockerclient.ContainerConfig{
			Image:"pfservice",
			Labels:labels,
			ExposedPorts:map[string]struct{}{
				"3000/tcp":{},
			},
		}

		var err error
		containerId, err = docker.CreateContainer(&containerConfig, "foobar")
		if err != nil {
			log.Fatal(err)
		}

		pb := dockerclient.PortBinding{HostPort:"8080"}
		portBindings := []dockerclient.PortBinding{pb}
		hostConfig := &dockerclient.HostConfig{
			PortBindings:map[string][]dockerclient.PortBinding {
				"3000/tcp":portBindings,
			},
		}
		err = docker.StartContainer(containerId, hostConfig)
		if err != nil {
			log.Fatal(err)
		}

		println("...we just started a docker container...")



	})

	Given(`^A prime factor resource value of (\d+)$`, func(i1 int) {
	})

	When(`^I call the prime factors service$`, func() {
	})

	Then(`^The prime factors for the resouce value are returned$`, func() {
	})

	After("@primefactors", func() {
		if containerId != "" {
			println("stop container")
			docker.StopContainer(containerId, 5)
		}
	})

}