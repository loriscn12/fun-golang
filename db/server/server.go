package main

import (
	"google.golang.org/grpc"

	"context"
	"flag"
	"fmt"
	"golang-project/db/mongodb"
	"golang-project/db/service"
	"log"
	"net"
	"time"

	dpb "golang-project/db/proto"
)

var (
	port      = flag.Int("port", 10000, "The server port")
	mongoAddr = flag.String("mongoAddress", "mongodb://localhost:27017", "The address of the MongoDB")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	client, err := mongodb.New(ctx, *mongoAddr)
	if err != nil {
		log.Fatalf("failed to instanciate a new MongoDB client: %s", err)
	}
	dbServer, err := service.New(ctx, &service.Config{MongoClient: client})
	if err != nil {
		log.Fatalf("failed to instanciate a new DBService client: %s", err)
	}
	defer dbServer.Close(ctx)
	grpcServer := grpc.NewServer()

	// TODO(loris): add MongoDB as a critical dependency for the gRPC server.
	dpb.RegisterDatabaseServer(grpcServer, dbServer)
	grpcServer.Serve(lis)
}
