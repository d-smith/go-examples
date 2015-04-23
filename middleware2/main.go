package main

import (
	"net/http"
)

type Wrapper interface {
	Wrap(http.Handler) http.Handler
}

type WrapperFactory interface {
	NewWrapper() Wrapper
}

type AWrapper struct{}

func (aw AWrapper) Wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		w.Write([]byte("A wrapper wrote this\n"))
	})

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

func chainWrappers(hf func(w http.ResponseWriter, r *http.Request), wrappers ...Wrapper) http.Handler {
	handler := http.HandlerFunc(hf)
	for _, wrapper := range wrappers {
		if wrapper == nil {
			continue
		}
		handler = (wrapper.Wrap(handler)).(http.HandlerFunc)
	}

	return handler
}

func main() {
	//Todo - update sample to work with slice of factories
	http.ListenAndServe(":8080", chainWrappers(handleCall, new(AWrapper), new(BWrapper)))
}
