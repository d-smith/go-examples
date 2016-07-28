package main

import (
	"errors"
	"log"
	"net/http"
)

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("ping\n"))
}

func pongHandler(rw http.ResponseWriter, req *http.Request) {
	log.Panic("ah, snap")
}

type RecoveryContext struct {
	logFn    func(interface{})
	errMsgFn func(interface{}) string
}

func (rc *RecoveryContext) MakeRecoveryWrapper(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				rc.logFn(r)
				http.Error(w, rc.errMsgFn(r), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	rc := RecoveryContext{
		logFn: func(r interface{}) {
			var err error
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = errors.New("Unknown error")
			}
			log.Println("Something went wrong", err.Error())
		},
		errMsgFn: func(r interface{}) string {
			return "Something went wrong - sorry about that"
		},
	}

	mux.Handle("/ping", rc.MakeRecoveryWrapper(pingHandler))
	mux.Handle("/pong", rc.MakeRecoveryWrapper(pongHandler))

	http.ListenAndServe(":4000", mux)
}
