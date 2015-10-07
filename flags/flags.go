package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type servers []string

func (s *servers) Set(val string) error {
	for _, v := range strings.Split(val, ",") {
		trimmed := strings.TrimSpace(v)
		*s = append(*s, trimmed)
	}
	return nil
}

func (s *servers) String() string {
	return fmt.Sprint(*s)
}

func main() {
	var cmdServers servers
	flag.Var(&cmdServers, "servers", "comma-separated list of servers")

	var checkInterval time.Duration
	flag.DurationVar(&checkInterval, "check-interval", 10*time.Second, "Interval between health checks")
	flag.Parse()
	fmt.Println(cmdServers)

	fmt.Printf("Now plus check-interval: %v", time.Now().Add(checkInterval))
}
