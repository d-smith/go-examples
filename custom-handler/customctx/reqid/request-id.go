package reqid

import (
	cc "github.com/d-smith/go-examples/custom-handler/customctx"
	"golang.org/x/net/context"
	"net/http"
)

type key int

const requestIDKey key = 0

func NewContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, requestIDKey, req.Header.Get("X-Request-ID"))
}

func RequestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

func RequestIdMiddleware(h cc.ContextHandler) cc.ContextHandler {
	return cc.ContextHandlerFunc(func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		ctx = NewContextWithRequestID(ctx, req)
		println("-->Request id serve http")
		h.ServeHTTPContext(ctx, rw, req)
		println("<--Request id http served")
	})
}
