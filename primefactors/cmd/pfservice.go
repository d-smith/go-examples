package main

import (
	"encoding/json"
	_ "expvar"
	"flag"
	"fmt"
	"github.com/d-smith/go-examples/primefactors"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func extractResource(uri string) (string, error) {
	parts := strings.Split(uri, "/")
	if len(parts) != 3 || parts[2] == "" {
		return "", fmt.Errorf("Expected URI format: /foo/<resource id>")
	}

	return parts[2], nil

}

func handleCall(w http.ResponseWriter, r *http.Request) {
	log.Println("Go routines: ", runtime.NumGoroutine())
	//pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)

	resourceId, err := extractResource(r.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	n, err := strconv.Atoi(resourceId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Resource needs to be a positive integer"))
		return
	}

	log.Println("Calculate prime factors of ", n)
	pf := primefactors.PrimeFactors(n)

	bytes, _ := json.Marshal(pf)
	w.Write(bytes)

}

func main() {
	var port = flag.Int("port", -1, "Port to listen on")
	flag.Parse()
	if *port == -1 {
		fmt.Println("Must specify a -port argument")
		return
	}

	http.Handle("/pf/", http.HandlerFunc(handleCall))
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
