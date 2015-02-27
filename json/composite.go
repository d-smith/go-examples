package main

import (
	"encoding/json"
	"fmt"
)

type managedService struct {
	Address      string
	ListenerName string
	Routes       []route
}

type route struct {
	Name    string
	URIRoot string
	Backend backend
}

type backend struct {
	Name    string
	Servers []server
}

type server struct {
	Name    string
	Address string
	PingURI string
	Up      bool
}

func compositeMain() {
	jsonDoc :=
	`
	{
		"address": "0.0.0.0:8080",
		"listenerName": "l1",
		"routes":[{
			"name": "route1",
			"uirRoot": "/foo",
			"backend": {
				"name":"be1",
				"servers":[
				{
					"name":"s1",
					"address":"localhost:3000",
					"pingUri":"/hello",
					"up":true
				}
				]
			}
			
		}]
	}
	`
	ms := new(managedService)
	json.Unmarshal([]byte(jsonDoc),ms)
	
	fmt.Println(ms)
}