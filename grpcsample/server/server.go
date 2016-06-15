package main

import (
	"flag"
	"fmt"
	. "github.com/d-smith/go-examples/grpcsample"
	"github.com/mwitkow/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"math/rand"
	"net"
	"time"
)

type MyDirectoryServer struct{}

func LogRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	grpclog.Printf("Request:%v", req)
	return handler(ctx, req)
}

func LogDestination(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	grpclog.Printf("Dest:%v", info.FullMethod)
	return handler(ctx, req)
}

func (mds *MyDirectoryServer) LookupPersonByName(context.Context, *NameRequest) (*Person, error) {
	grpclog.Println("Receieved request... sleep")
	time.Sleep(randomSleepTime())
	grpclog.Print("Ok - found it")
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

func randomSleepTime() time.Duration {
	millis := rand.Intn(1000) + 1
	delay := time.Duration(millis) * time.Millisecond
	return delay
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	//var opts []grpc.ServerOption
	var grpcServer *grpc.Server

	grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(LogRequest, LogDestination)))
	RegisterDirectoryServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
