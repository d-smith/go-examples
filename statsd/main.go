package main

import (
	"time"
	"github.com/quipo/statsd"
	"math/rand"
)

func init() {
	rand.Seed(123123213)
}

const (
	maxTime = 100
	minTime = 10
)

func main() {
	prefix := "myproject."
    statsdclient := statsd.NewStatsdClient("172.20.20.10:8125", prefix)
    statsdclient.CreateSocket()
    interval := time.Second * 2 // aggregate stats and flush every 2 seconds
    stats := statsd.NewStatsdBuffer(interval, statsdclient)
    defer stats.Close()

    // not buffered: send immediately
    for {
	    statsdclient.Incr("mymetric", 4)
	
	    // buffered: aggregate in memory before flushing
	    stats.Incr("mymetric", 1)
	    stats.Incr("mymetric", 3)
	    stats.Incr("mymetric", 1)
	    stats.Incr("mymetric", 1)
	    
	    
	    
	    stats.Gauge("circuit-breaker-1", 0)
	    
	    time.Sleep(300 * time.Millisecond)
	    stats.Timing("loop.time", rand.Int63n(maxTime- minTime) + minTime)
    }
}

