package main

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"io/ioutil"
	"log"
	"net/http"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

const HealthyEndpointsOnly = true

func main() {

	log.Println("create client")
	config := &consul.Config{
		Address:    "172.20.20.70:8500",
		Scheme:     "http",
		HttpClient: http.DefaultClient,
	}

	client, _ := consul.NewClient(config)

	log.Println("look for the demo service in envdev1")
	catalogServices, _, err := client.Catalog().Service("demo-service", "envdev1", nil)
	fatal(err)

	log.Println("demo service endpoints")
	for _, svc := range catalogServices {
		fmt.Printf("%s:%d\n", svc.ServiceAddress, svc.ServicePort)
	}

	healthChecks, _, err := client.Health().Service("demo-service", "", HealthyEndpointsOnly, nil)
	fatal(err)

	log.Println("invoke healty endpoints only")
	for _, hc := range healthChecks {
		endpoint := fmt.Sprintf("http://%s:%d/hello", hc.Service.Address, hc.Service.Port)
		fmt.Printf("Endpoint: %s\n", endpoint)
		resp, err := http.Get(endpoint)
		fatal(err)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fatal(err)
		fmt.Printf("Response:\n%s\n", body)
	}
}
