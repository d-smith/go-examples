package main

import (
	"github.com/armon/go-metrics"
	"github.com/armon/go-metrics/datadog"
	"log"
	"net/http"
	"time"
)

type apiHandler struct{}

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		defer metrics.MeasureSince([]string{"FooService"}, time.Now())
		log.Println("request to foo")
		w.Write([]byte("uh-huh\n"))
	})

	/*
		inm := metrics.NewInmemSink(10*time.Second, time.Minute)
		metrics.DefaultInmemSignal(inm)
		metrics.NewGlobal(metrics.DefaultConfig("service-name"), inm)
	*/

	sink, err := datadog.NewDogStatsdSink("localhost:8125", "devmac.xtrac")

	//	sink, err := metrics.NewStatsdSink("localhost:8125")
	if err != nil {
		log.Fatal(err)
	}
	metrics.NewGlobal(metrics.DefaultConfig("xavi"), sink)

	http.ListenAndServe(":8080", nil)
}
