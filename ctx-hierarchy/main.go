package main

import (
	"github.com/d-smith/go-examples/ctx-hierarchy/service"
	"github.com/d-smith/go-examples/ctx-hierarchy/wrappers"
	"github.com/d-smith/go-examples/custom-handler/customctx"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func main() {

	ctxHandlerFn := customctx.ContextHandlerFunc(service.NewBroHandler())
	wrapper := wrappers.RequestTimerMiddleware(wrappers.RequestIdMiddleware(ctxHandlerFn))
	h := &customctx.ContextAdapter{
		Ctx:     context.Background(),
		Handler: wrapper,
	}
	log.Println("starting listener on port 8080")
	http.ListenAndServe(":8080", h)
}
