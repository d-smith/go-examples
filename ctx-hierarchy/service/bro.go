package service

import (
	"golang.org/x/net/context"
	"net/http"
	"sync"
)

//Yes - this is inspired by the bro app on silicon valley

func NewBroHandler() func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		var wg sync.WaitGroup
		wg.Add(2)

		ctxA := context.WithValue(ctx, 100, "service-x")
		ctxB := context.WithValue(ctx, 100, "service-y")

		go func(ctx context.Context) {
			defer wg.Done()
			println((ctx.Value(100)).(string))
			bro(ctx,rw)
		}(ctxA)

		go func(ctx context.Context) {
			defer wg.Done()
			println((ctx.Value(100)).(string))
			bro(ctx,rw)
		}(ctxB)
		wg.Wait()

	}
}


func bro(ctx context.Context, rw http.ResponseWriter) {
	rw.Write([]byte("Bro\n"))
}