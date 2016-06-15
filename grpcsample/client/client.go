package main

import (
	"flag"
	. "github.com/d-smith/go-examples/grpcsample"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"time"
)

var (
	serverAddr = flag.String("serveraddr", "127.0.0.1:10000", "The server address in the format of host:port")
	doTimeout  = flag.Bool("dotimeout", false, "set to true for forcing client timeout")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	nameReq := new(NameRequest)
	nameReq.Name = "flibby"

	client := NewDirectoryClient(conn)

	ctx := context.Background()
	var cancel context.CancelFunc
	if *doTimeout == true {
		ctx, cancel = context.WithTimeout(ctx, 500*time.Millisecond)
		defer cancel()
	}

	person, err := client.LookupPersonByName(ctx, nameReq)
	if err != nil {
		grpclog.Fatalf("error calling remote service: %v", err)
	}

	grpclog.Print(person)

}
