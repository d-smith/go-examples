package main

import (
	"github.com/hashicorp/consul/api"
	"log"
)

func main() {
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	catalog := client.Catalog()
	env1Services, _, err := catalog.Service("demo-service", "env1", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	env1Service := env1Services[0]

	log.Printf("Call env1 demo-service at %s:%d\n", env1Service.ServiceAddress, env1Service.ServicePort)

	log.Println("bye for now")
}
