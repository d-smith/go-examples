package main

import (
	"flag"
	"fmt"
	. "github.com/d-smith/go-examples/grpcsample"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type MyDirectoryServer struct{}

func (mds *MyDirectoryServer) LookupPersonByName(context.Context, *NameRequest) (*Person, error) {
	return &Person{
		Name:  "flibby",
		Email: "flibby@dibby-dibby-do.com",
	}, nil
}

var (
	port = flag.Int("port", 10000, "Server port")
)

func newServer() *MyDirectoryServer {
	s := new(MyDirectoryServer)
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	var grpcServer *grpc.Server

	grpcServer = grpc.NewServer(opts...)
	RegisterDirectoryServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
