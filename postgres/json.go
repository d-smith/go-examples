package main

type managedService struct {
	XMLName      struct{} `xml:"xmlManagedService"`
	Address      string   `json:"address" xml:"xmlAddress"`
	ListenerName string   `json:"listenerName" xml:"xmlListenerName"`
	Routes       []route  `json:"routes" xml:"xmlRoutes"`
}

type route struct {
	Name    string  `json:"name" xml:"xmlName"`
	URIRoot string  `json:"uriRoot" xml:"xmlUriRoot"`
	Backend backend `json:"backend" xml:"xmlBackend"`
}

type backend struct {
	Name    string   `json:"name" xml:"xmlName"`
	Servers []server `json:"servers" xml:"xmlServers"`
}

type server struct {
	Name    string `json:"name" xml:"xmlName"`
	Address string `json:"address" xml:"xmlAddress"`
	PingURI string `json:"pingUri" xml:"xmlPingUri"`
	Up      bool   `json:"up" xml:"xmlUp"`
}

var jsonDoc =
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
		},
		{
			"name":"s2",
			"address":"localhost:3100",
			"pingUri":"/hello",
			"up":true
		}
		]
	}

}]
}
`