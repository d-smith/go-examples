package main

import (
	"github.com/d-smith/go-examples/ctx-hierarchy/service"
	"github.com/d-smith/go-examples/custom-handler/customctx"
	"golang.org/x/net/context"
	"net/http"
	"github.com/d-smith/go-examples/ctx-hierarchy/wrappers"
)

func main() {

	ctxHandlerFn := customctx.ContextHandlerFunc(service.NewBroHandler())
	wrapper := wrappers.RequestIdMiddleware(ctxHandlerFn)
	h := &customctx.ContextAdapter{
		Ctx:     context.Background(),
		Handler: wrapper,
	}
	http.ListenAndServe(":8080", h)
}
