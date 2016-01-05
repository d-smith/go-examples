package service

import (
	"golang.org/x/net/context"
	"net/http"
)

//Yes - this is inspired by the bro app on silicon valley

func NewBroHandler() func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Bro\n"))
	}
}
