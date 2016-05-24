package main

import (
	"os"
	"log"
	"strconv"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"io"
	"github.com/hashicorp/consul/api"
)

var transport = &http.Transport{DisableKeepAlives: false, DisableCompression: false}

//Note this assumes we're running in docker and that the actual hostname has
//been bound to the logical consul hostname via the --link command line option.
func lookupServiceEndpointOrDie(env string)string {
	config := api.DefaultConfig()
	config.Address = "consul:8500"
	client,err := api.NewClient(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	catalog := client.Catalog()
	log.Println("Lookup demo-service for environment",env)
	env1Services,_,err := catalog.Service("demo-service",env,nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	if len(env1Services) == 0 {
		log.Fatal("No service definition found in consul")
	}

	env1Service := env1Services[0]

	return fmt.Sprintf("%s:%d", env1Service.ServiceAddress, env1Service.ServicePort)
}

func getConfigFromEnvOrDieTrying()(endpoint string, port int) {
	env := os.Getenv("env")
	if env == "" {
		log.Fatal("No environment designation specified via the env envronment variable")
	}

	endpoint = lookupServiceEndpointOrDie(env)

	port,err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		log.Fatal("No valid port specified")
	}

	return
}

func makeHandler(endpoint string) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = endpoint
		req.Host = endpoint
		resp, err := transport.RoundTrip(req)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		io.Copy(rw, resp.Body)
		resp.Body.Close()
	}
}

func main() {
	endpoint, port := getConfigFromEnvOrDieTrying()
	log.Println("Port:", port, "Endpoint:",endpoint)

	r := mux.NewRouter()
	r.HandleFunc("/",makeHandler(endpoint))
	http.Handle("/",r)
	err := http.ListenAndServe(fmt.Sprintf(":%d",port), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
