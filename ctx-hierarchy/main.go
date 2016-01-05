package main

import (
	"github.com/d-smith/go-examples/ctx-hierarchy/service"
	"github.com/d-smith/go-examples/custom-handler/customctx"
	"golang.org/x/net/context"
	"net/http"
)

func main() {

	h := &customctx.ContextAdapter{
		Ctx:     context.Background(),
		Handler: customctx.ContextHandlerFunc(service.NewBroHandler()),
	}
	http.ListenAndServe(":8080", h)
}
