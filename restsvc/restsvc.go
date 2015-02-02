/*
Worked through the REST example from 
http://dougblack.io/words/a-restful-micro-framework-in-go.html

1st change - need a way to default a method not supported response instead
of declaring not supported methods for each HTTP method that each resource type
must include. Ok, I peeked at sleepy... seriously go here for the good stuff:

https://github.com/dougblack/sleepy

This code is a learning exercise.
*/


package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
	"log"
)

func response(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("Hello world."))
}

type GetSupported interface {
	Get(url.Values, http.Header) (int, interface{}, http.Header)
}

type PostSupported interface {
	Post(url.Values, http.Header) (int, interface{}, http.Header)
}



//Receiver for methods that manage our resources
type API struct {} 


//TODO - can I use a type switch here? (see Effective Go)
func (api *API) requestHandler(resource interface{}) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		
        
		request.ParseForm()  //TODO: error handling
		
		var handler func(url.Values, http.Header) (int, interface{}, http.Header)
		
		switch request.Method {
			case "GET":
				if resource, ok := resource.(GetSupported); ok {
					handler = resource.Get	
				}
			case "POST":
				if resource, ok := resource.(PostSupported); ok {
					handler = resource.Post	
				} 
		}
		
		
		if handler == nil {
			rw.WriteHeader(500)
			return
		}
		
		code, data, header := handler(request.Form, request.Header)
		
		content, err := json.Marshal(data)
		if err != nil {
			rw.WriteHeader(500)
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
    return 200, data, http.Header{"Content-type": {"application/json"}}
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