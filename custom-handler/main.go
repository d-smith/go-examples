package main

/*
	From https://joeshaw.org/net-context-and-http-handler/
 */

import (
	"golang.org/x/net/context"
	"net/http"
	"fmt"
	"github.com/d-smith/go-examples/custom-handler/rc"
)





func handler(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	reqID := rc.RequestIDFromContext(ctx)
	fmt.Fprintf(rw, "Hello request ID %s\n", reqID)
}



func main() {
	h := &rc.ContextAdapter{
		Ctx: context.Background(),
		Handler: rc.Middleware(rc.ContextHandlerFunc(handler)),
	}
	http.ListenAndServe(":8080", h)
}

