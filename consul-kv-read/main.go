package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/consul/api"
)

// List some keys prepopulated in consul via the admin console.
func main() {
	client, _ := api.NewClient(api.DefaultConfig())
	kv := client.KV()

	pairs, _, err := kv.List("/listeners/", nil)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%v pairs returned\n", len(pairs))

	for i := range pairs {
		pair := pairs[i]
		fmt.Printf("key %s value %s\n", pair.Key, pair.Value)
	}
}
