package main

import (
	"net"
	"net/rpc"
	"time"
)

type (
	PingClient struct {
		connection *rpc.Client
	}
)

func NewPingClient(dsn string, timeout time.Duration) (*PingClient, error) {
	connection, err := net.DialTimeout("tcp", dsn, timeout)
	if err != nil {
		return nil, err
	}
	return &PingClient{connection: rpc.NewClient(connection)}, nil
}

func (pc *PingClient) Ping(in string) (string, error) {
	var out *PingResponse
	err := pc.connection.Call("RPC.Ping", in, &out)
	return out.Ping, err
}
