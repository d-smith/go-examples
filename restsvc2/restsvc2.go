package main

import (
	"encoding/json"
	"expvar"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

var (
	counts = expvar.NewMap("counters")
)

type HttpHandler func(http.ResponseWriter, *http.Request)

type GetSupported interface {
	Get(http.ResponseWriter, *http.Request)
}

type PostSupported interface {
	Post(http.ResponseWriter, *http.Request)
}

type API struct{}

func getHandlerForMethod(method string, resource interface{}) (handler HttpHandler) {
	switch method {
	case GET:
		if resource, ok := resource.(GetSupported); ok {
			handler = resource.Get
		}
	case POST:
		if resource, ok := resource.(PostSupported); ok {
			handler = resource.Post
		}
	}
	return
}

func writeResponse(code int, content []byte, header http.Header, rw http.ResponseWriter) {
	for name, values := range header {
		for _, value := range values {
			rw.Header().Add(name, value)
		}
	}
	rw.WriteHeader(code)
	rw.Write(content)
}

func (api *API) requestHandler(resource interface{}) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {

		handler := getHandlerForMethod(request.Method, resource)
		if handler == nil {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		handler(rw, request)

	}
}

func counterName(method string, path string) string {
	return fmt.Sprintf("%s::%s", GET, path)
}

func incCounter(method string, path string) {
	//TODO - design this to handle concurrency
	//...and how to increment	
}

func initCountersForResource(path string) {
	counts.Add(counterName(GET, path), 0)
	counts.Add(counterName(POST, path), 0)
}

func (api *API) AddResource(resource interface{}, path string) {
	initCountersForResource(path)
	http.HandleFunc(path, api.requestHandler(resource))
}

func (api *API) Start(port int) {
	portString := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type HelloResource struct{}

const (
	helloGets  = "HelloGets"
	helloPosts = "HelloPosts"
)

func (HelloResource) Get(rw http.ResponseWriter, req *http.Request) {
	data := map[string]string{"hello": "world"}
	header := http.Header{"Content-type": {"application/json"}}

	content, err := json.Marshal(data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeResponse(http.StatusOK, content, header, rw)
}

type HelloObj struct {
	X int
	Y int
}

func (HelloResource) Post(rw http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	request.Body.Close()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	var hello HelloObj
	err = json.Unmarshal(body, &hello)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	log.Println(hello)

	writeResponse(http.StatusOK, []byte("Good to go"), nil, rw)

}

func serverMonitorInterface() {
	sock, err := net.Listen("tcp", "localhost:8123")
	if err != nil {
		panic(err)
	}
	go func() {
		http.Serve(sock, nil)
	}()
}

func main() {
	// start a monitor another port
	serverMonitorInterface()

	helloResource := new(HelloResource)

	var api = new(API)
	api.AddResource(helloResource, "/hello")
	log.Println("...starting...")
	api.Start(3000)
}
