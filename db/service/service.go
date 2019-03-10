package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	"log"
	"fmt"
	"net"
	"context"
	"flag"

	dpb "golang-project/db/proto"
	grpc "google.golang.org/grpc"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type dbServer struct {
	dbClient *mongo.Client
}

// Test returns a test response.
func (s *dbServer) AddUser(ctx context.Context, req *dpb.AddUserRequest) (*dpb.AddUserResponse, error) {
	err := s.dbClient.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatalf("failed to connect to MongoDB client: %s", err)
	}
	collection := s.dbClient.Database("test").Collection("users")
	_, err = collection.InsertOne(ctx, bson.M{"name": req.Name, "surname": req.Surname})
	if err != nil{
		log.Fatalf("failed to insert to MongoDB: %s", err)
	}
	return &dpb.AddUserResponse{Id: "test"}, nil
}

func main() {
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel();
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed to instantiate a new MongoDB client: %s", err)
	}
	dpb.RegisterDatabaseServer(grpcServer, &dbServer{dbClient: client})
	grpcServer.Serve(lis)
}
