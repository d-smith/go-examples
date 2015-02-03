package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

type HttpHandler func(*http.Request) (int, interface{}, http.Header)

type GetSupported interface {
	Get(*http.Request) (int, interface{}, http.Header)
}

type PostSupported interface {
	Post(*http.Request) (int, interface{}, http.Header)
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

		code, data, header := handler(request)

		content, err := json.Marshal(data)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		writeResponse(code, content, header, rw)

	}
}

func (api *API) AddResource(resource interface{}, path string) {
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

func (HelloResource) Get(*http.Request) (int, interface{}, http.Header) {
	data := map[string]string{"hello": "world"}
	return http.StatusOK, data, http.Header{"Content-type": {"application/json"}}
}

type HelloObj struct {
	X int
	Y int
}

func (HelloResource) Post(request *http.Request) (int, interface{}, http.Header) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, nil
	}

	var hello HelloObj
	err = json.Unmarshal(body, &hello)
	if err != nil {
		return http.StatusInternalServerError, nil, nil
	}

	log.Println(hello	)

	return http.StatusOK, "Good to go", nil
}

func main() {
	helloResource := new(HelloResource)

	var api = new(API)
	api.AddResource(helloResource, "/hello")
	log.Println("...starting...")
	api.Start(3000)
}
