package wrappers
import (
	"golang.org/x/net/context"
	"net/http"
	cc "github.com/d-smith/go-examples/custom-handler/customctx"
	"log"
)

type key int

const idkey key = 1

func NewContextWithBroID(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, idkey, req.Header.Get("X-Bro-ID"))
}

func RequestIDFromContext(ctx context.Context) string {
	return ctx.Value(idkey).(string)
}

func RequestIdMiddleware(h cc.ContextHandler) cc.ContextHandler {
	return cc.ContextHandlerFunc(func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		ctx = NewContextWithBroID(ctx, req)
		log.Println("-->Put Bro into context:", RequestIDFromContext(ctx))
		h.ServeHTTPContext(ctx, rw, req)
	})
}