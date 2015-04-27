package main

import (
	"net/http"
)

//Wrapper defines an interface for things that can wrap http Handlers
type Wrapper interface {
	Wrap(http.Handler) http.Handler
}

//WrapperFactory defines a function that can create something that
//implements Wrapper
type WrapperFactory func() Wrapper

//NewAWrapper instantiates AWrapper
func NewAWrapper() Wrapper {
	return new(AWrapper)
}

//AWrapper can wrap http handlers
type AWrapper struct{}

//Wrap wraps http.Handlers with A stuff
func (aw AWrapper) Wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		w.Write([]byte("A wrapper wrote this\n"))
	})
}

//NewBWrapper instantiates BWrappers
func NewBWrapper() Wrapper {
	return new(BWrapper)
}

//BWrapper can wrap http handlers
type BWrapper struct{}

//Wrap wraps am http Handler adding BWrapperness to the wrapped handler
func (bw BWrapper) Wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		w.Write([]byte("B wrapper wrote this\n"))
	})

}

func handleCall(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handleCall wrote this stuff\n"))
}

func chainWrappers(hf func(w http.ResponseWriter, r *http.Request), wrapperFactories []WrapperFactory) http.Handler {
	handler := http.HandlerFunc(hf)
	for _, factory := range wrapperFactories {
		if factory == nil {
			continue
		}
		wrapper := factory()
		handler = (wrapper.Wrap(handler)).(http.HandlerFunc)
	}

	return handler
}

func main() {
	factories := []WrapperFactory{NewAWrapper, NewBWrapper}
	http.ListenAndServe(":8080", chainWrappers(handleCall, factories))
}
