package main

import (
	"net/http"
)

type Wrapper interface {
	Wrap(http.Handler) http.Handler
}

type WrapperFactory func() Wrapper

func NewAWrapper() Wrapper {
	return new(AWrapper)
}

type AWrapper struct{}

func (aw AWrapper) Wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		w.Write([]byte("A wrapper wrote this\n"))
	})

}

func NewBWrapper() Wrapper {
	return new(BWrapper)
}

type BWrapper struct{}

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
	factories := []WrapperFactory{NewAWrapper,NewBWrapper}
	http.ListenAndServe(":8080", chainWrappers(handleCall, factories))
}
