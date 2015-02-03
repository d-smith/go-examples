/*
Worked through the REST example from 
http://dougblack.io/words/a-restful-micro-framework-in-go.html

1st change - need a way to default a method not supported response instead
of declaring not supported methods for each HTTP method that each resource type
must include. Ok, I peeked at sleepy... seriously go here for the good stuff:

https://github.com/dougblack/sleepy

This code is a learning exercise.

2nd experiment - played with an alternative way of getting the handler method,
by switching on the resource type:

switch t := resource.(type) {
	case GetSupported:
		if request.Method == "GET" {
			handler = t.Get
		}
	case PostSupported:
		if(request.Method == "POST") {
			handler = t.Post
		}
}
		
This also works, but I think the code reads better by first considering the HTTP method,
then determining if the resource can handle the method.		

Next refactors - Use of http status codes, not constants.

*/


package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
	"log"
)

const (
	GET = "GET"
	POST = "POST"
)


type GetSupported interface {
	Get(url.Values, http.Header) (int, interface{}, http.Header)
}

type PostSupported interface {
	Post(url.Values, http.Header) (int, interface{}, http.Header)
}



//Receiver for methods that manage our resources
type API struct {} 


func (api *API) requestHandler(resource interface{}) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		
        
		if request.ParseForm() != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return	
		}
		
		var handler func(url.Values, http.Header) (int, interface{}, http.Header)
		
		switch request.Method {
			case GET:
				if resource, ok := resource.(GetSupported); ok {
					handler = resource.Get	
				}
			case POST:
				if resource, ok := resource.(PostSupported); ok {
					handler = resource.Post	
				} 
		}
		
		
		if handler == nil {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		
		code, data, header := handler(request.Form, request.Header)
		
		content, err := json.Marshal(data)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		for name, values := range header {
			for _, value := range values {
				rw.Header().Add(name, value)
			}
		}
		rw.WriteHeader(code)
		rw.Write(content)
		
	}
}

func (api *API) AddResource(resource interface{}, path string) {
    http.HandleFunc(path, api.requestHandler(resource))
}

type HelloResource struct {}

func (HelloResource) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
    data := map[string]string{"hello": "world"}
    return http.StatusOK, data, http.Header{"Content-type": {"application/json"}}
}

func (HelloResource) Post(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	for k:= range values {
		fmt.Println("key: ", k, " value ", values[k])
	}
	return http.StatusOK, nil, nil
}

func (api *API) Start(port int) {
    portString := fmt.Sprintf(":%d", port)
    err := http.ListenAndServe(portString, nil)
    if err != nil {
    	log.Fatal("ListenAndServe: ", err)
    }
}

func main() {
	
	helloResource := new(HelloResource)
	
	var api = new(API)
	api.AddResource(helloResource, "/hello")
	fmt.Println("...starting...")
	api.Start(3000)
	
}