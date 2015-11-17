package main

/*
	From https://joeshaw.org/net-context-and-http-handler/
*/

import (
	"fmt"
	"github.com/d-smith/go-examples/custom-handler/customctx"
	"github.com/d-smith/go-examples/custom-handler/customctx/reqid"
	"golang.org/x/net/context"
	"net/http"
)

func handler(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	reqID := reqid.RequestIDFromContext(ctx)
	fmt.Fprintf(rw, "Hello request ID %s\n", reqID)
}

func main() {
	h := &customctx.ContextAdapter{
		Ctx:     context.Background(),
		Handler: reqid.Middleware(customctx.ContextHandlerFunc(handler)),
	}
	http.ListenAndServe(":8080", h)
}
