package wrappers

import (
	cc "github.com/d-smith/go-examples/custom-handler/customctx"
	"golang.org/x/net/context"
	"net/http"
)

func NewContextWithBroID(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, idkey, req.Header.Get("X-Bro-ID"))
}

func RequestIDFromContext(ctx context.Context) string {
	return ctx.Value(idkey).(string)
}

func RequestIdMiddleware(h cc.ContextHandler) cc.ContextHandler {
	return cc.ContextHandlerFunc(func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		ctx = NewContextWithBroID(ctx, req)
		h.ServeHTTPContext(ctx, rw, req)
	})
}
