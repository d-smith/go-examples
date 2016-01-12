package wrappers

import (
	cc "github.com/d-smith/go-examples/custom-handler/customctx"
	"github.com/d-smith/go-examples/timer"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func NewContextWithTimer(ctx context.Context, req *http.Request) context.Context {
	timer := timer.NewAPITimer("generic timer")
	return context.WithValue(ctx, timerKey, timer)
}

func TimerFromContext(ctx context.Context) *timer.APITime {
	newCtx := ctx.Value(timerKey).(*timer.APITime)
	return newCtx
}

func RequestTimerMiddleware(h cc.ContextHandler) cc.ContextHandler {
	return cc.ContextHandlerFunc(func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		ctx = NewContextWithTimer(ctx, req)
		h.ServeHTTPContext(ctx, rw, req)
		ctxTimer := TimerFromContext(ctx)
		ctxTimer.Stop(nil)
		go func(t *timer.APITime) {
			log.Println(t.ToJSONString())
		}(ctxTimer)
	})
}
