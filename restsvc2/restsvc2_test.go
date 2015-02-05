package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"fmt"
)

const (
	expectedGetResponse = "here's your response"
)

type TestResource struct {}

func (TestResource) Get(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(expectedGetResponse))
}

func testGetHandling(getHandler HttpHandler, t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getHandler))
	defer ts.Close()
	
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("Error returned on GET", err)
	}
	
	rs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()	
	if err != nil {
		t.Error("Error returned reading response", err)
	}	
	
	responseString := string(rs)
	if responseString != expectedGetResponse {
		t.Error(fmt.Sprintf("expected %s got %s", expectedGetResponse, responseString))
	}
}


func TestMethodHandlerExtraction(t *testing.T) {
	
	testResource := new(TestResource)
	
	getHandler := getHandlerForMethod("GET", testResource)
	if getHandler == nil {
		t.Error("Nil handler returned for interface defining GET handler")	
	}
	
	testGetHandling(getHandler,t)
	
	postHandler := getHandlerForMethod("POST", testResource)
	if postHandler != nil {
		t.Error("Non-nil handler returned for interface that doesn't handle POST")	
	}
}