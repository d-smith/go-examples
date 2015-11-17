package reqid

import (
	"fmt"
	cc "github.com/d-smith/go-examples/custom-handler/customctx"
	"golang.org/x/net/context"
	"net/http"
)

type key int

const requestIDKey key = 0

func NewContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	fmt.Println(req.Header)
	return context.WithValue(ctx, requestIDKey, req.Header.Get("X-Request-ID"))
}

func RequestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

func Middleware(h cc.ContextHandler) cc.ContextHandler {
	return cc.ContextHandlerFunc(func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		ctx = NewContextWithRequestID(ctx, req)
		h.ServeHTTPContext(ctx, rw, req)
	})
}
