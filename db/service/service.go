package main

import (
	"log"
	"fmt"
	"net"
	"context"
	"flag"

	dpb "golang-project/db/proto"
	grpc "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type dbServer struct {
}

// Test returns a test response.
func (s *dbServer) Test(ctx context.Context, req *dpb.TestRequest) (*dpb.TestResponse, error) {
	return &dpb.TestResponse{Test: "response"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	dpb.RegisterDatabaseServer(grpcServer, &dbServer{})
	grpcServer.Serve(lis)
}
