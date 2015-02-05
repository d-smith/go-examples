package main

import (
	"testing"
	"net/http"
)

type TestResource struct {}

func (TestResource) Get(http.ResponseWriter, *http.Request) {}


func TestMethodHandlerExtraction(t *testing.T) {
	
	testResource := new(TestResource)
	
	getHandler := getHandlerForMethod("GET", testResource)
	if getHandler == nil {
		t.Error("Nil handler returned for interface defining GET handler")	
	}
	
	postHandler := getHandlerForMethod("POST", testResource)
	if postHandler != nil {
		t.Error("Non-nil handler returned for interface that doesn't handle POST")	
	}
}