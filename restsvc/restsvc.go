/*
Worked through the sleepy example from 
http://dougblack.io/words/a-restful-micro-framework-in-go.html
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

type Resource interface {
	Get(values url.Values)(int, interface{})
	Post(values url.Values)(int, interface{})
}


type (
	GetNotSupported struct{}
	PostNotSupported   struct{}
)

func (GetNotSupported) Get(values url.Values) (int, interface{}) {
    return 405, ""
}

func (PostNotSupported) Post(values url.Values) (int, interface{}) {
    return 405, ""
}

//Receiver for methods that manage our resources
type API struct {} 

func (api *API) Abort(rw http.ResponseWriter, statusCode int) {
    rw.WriteHeader(statusCode)
}

func (api *API) requestHandler(resource Resource) http.HandlerFunc {
	fmt.Println("request handler called")
	return func(rw http.ResponseWriter, request *http.Request) {
		
		var data interface{}
        var code int
        
		method := request.Method
		request.ParseForm()
		values := request.Form
		
		switch method {
			case "GET":
				code, data = resource.Get(values)
		}
		
		content, err := json.Marshal(data)
		if err != nil {
			api.Abort(rw, 500)
		}
		
		rw.WriteHeader(code)
		rw.Write(content)
		
	}
}

func (api *API) AddResource(resource Resource, path string) {
    http.HandleFunc(path, api.requestHandler(resource))
    fmt.Println("handle added")
}

type HelloResource struct {
	PostNotSupported
}

func (HelloResource) Get(values url.Values) (int, interface{}) {
    data := map[string]string{"hello": "world"}
    return 200, data
}

func (api *API) Start(port int) {
	fmt.Println("api Start called")
    portString := fmt.Sprintf(":%d", port)
    fmt.Printf("listen and serve on %s\n", portString)
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