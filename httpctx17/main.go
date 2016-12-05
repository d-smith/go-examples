package main

import (
	"net/http"
	"sync"
	"golang.org/x/net/context"
	"time"
	"math/rand"
	"fmt"
)

type sessionKey int

const SessionKey sessionKey = 111

var seed = rand.NewSource(time.Now().UnixNano())
var gen = rand.New(seed)
var mutex sync.Mutex

func WrapWithSessionId(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		mutex.Lock()
		val := gen.Intn(999999999)
		mutex.Unlock()

		newR := r.WithContext(context.WithValue(r.Context(), SessionKey, val))

		h.ServeHTTP(w, newR)
	})
}

func HelloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var msg string = "rats no context"
		ctx := r.Context()
		if ctx != nil {
			var sessionId = -1
			sessionId,ok := ctx.Value(SessionKey).(int)
			if ok {
				msg = fmt.Sprintf("Hello to session %d\n", sessionId)
			}
		}

		w.Write([]byte(msg))
	})
}


func main() {
	wrapped := WrapWithSessionId(HelloHandler())
	http.Handle("/foo", wrapped)
	http.ListenAndServe(":3000", nil)
}


