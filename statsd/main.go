package main

import (
	"time"
	"github.com/quipo/statsd"
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
    	startTime := time.Now()
	    statsdclient.Incr("mymetric", 4)
	
	    // buffered: aggregate in memory before flushing
	    stats.Incr("mymetric", 1)
	    stats.Incr("mymetric", 3)
	    stats.Incr("mymetric", 1)
	    stats.Incr("mymetric", 1)
	    
	    
	    
	    stats.Gauge("circuit-breaker-1", 0)
	    
	    time.Sleep(5 * time.Second)
	    stats.PrecisionTiming("loop.time", time.Since(startTime))
    }
}

