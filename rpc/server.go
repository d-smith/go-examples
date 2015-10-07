package main

import (
	"log"
	"net"
	"net/rpc"
	"runtime"
	"strings"
)

type RPC struct{}

type PingResponse struct {
	Ping string
}

func (rpc *RPC) Ping(in string, out *PingResponse) error {
	upper := strings.ToUpper(in)
	*out = PingResponse{upper}
	return nil
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	rpc.Register(new(RPC))

	l, e := net.Listen("tcp", ":9876")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	rpc.Accept(l)
}
