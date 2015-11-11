package main

import (
	"github.com/d-smith/go-metrics"
	"github.com/d-smith/go-metrics/datadog"
	"log"
	"net/http"
	"time"
)

type apiHandler struct{}

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		defer metrics.MeasureSince([]string{"Foo handler"}, time.Now())
		log.Println("request to foo")
		w.Write([]byte("uh-huh\n"))
	})

	/*
		inm := metrics.NewInmemSink(10*time.Second, time.Minute)
		metrics.DefaultInmemSignal(inm)
		metrics.NewGlobal(metrics.DefaultConfig("service-name"), inm)
	*/
	sink, err := datadog.NewDogStatsdSink("localhost:8125", "devmac.xtrac")
	if err != nil {
		log.Fatal(err)
	}
	metrics.NewGlobal(metrics.DefaultConfig("service-name"), sink)

	http.ListenAndServe(":8080", nil)
}
