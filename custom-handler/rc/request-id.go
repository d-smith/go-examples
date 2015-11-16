package rc

import (
	"golang.org/x/net/context"
	"net/http"
	"fmt"
)

func NewContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	fmt.Println(req.Header)
	return context.WithValue(ctx, requestIDKey, req.Header.Get("X-Request-ID"))
}

func RequestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}



func Middleware(h ContextHandler) ContextHandler {
	return ContextHandlerFunc(func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		ctx = NewContextWithRequestID(ctx, req)
		h.ServeHTTPContext(ctx, rw, req)
	})
}