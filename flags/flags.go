package main

import (
	"flag"
	"fmt"
	"strings"
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
	flag.Parse()
	fmt.Println(cmdServers)
}
