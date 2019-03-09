package main

import (
	"log"
	"fmt"
	"context"
	"flag"
	"time"

	dpb "golang-project/db/proto"
	grpc "google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := dpb.NewDatabaseClient(conn)

	// Looking for a valid feature
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	resp, err := client.Test(ctx, &dpb.TestRequest{Test: "test"})
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(resp)
}