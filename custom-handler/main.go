package main

/*
	From https://joeshaw.org/net-context-and-http-handler/
*/

import (
	"fmt"
	"github.com/d-smith/go-examples/custom-handler/customctx"
	"github.com/d-smith/go-examples/custom-handler/customctx/reqid"
	"github.com/d-smith/go-examples/custom-handler/customctx/timing"
	"github.com/d-smith/go-examples/custom-handler/services/quote"
	"golang.org/x/net/context"
	"net/http"
)

func handler(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	reqID := reqid.RequestIDFromContext(ctx)
	fmt.Fprintf(rw, "Hello request ID %s\n", reqID)
}

func main() {
	wrapped := quote.QuoteMiddleware(customctx.ContextHandlerFunc(quote.QuoteHandler))
	wrapped = timing.TimerMiddleware(wrapped)

	h := &customctx.ContextAdapter{
		Ctx:     context.Background(),
		Handler: wrapped,
	}
	http.ListenAndServe(":8080", h)
}
