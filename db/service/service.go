package service

import (
	"golang-project/db/mongodb"

	"google.golang.org/grpc/codes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc/status"

	"context"
	"log"

	dpb "golang-project/db/proto"
)

// For testing purposes.
var logFatalf = log.Fatalf

// DBService is the service's client.
type DBService struct {
	Client mongodb.Client
}

// Config contains config parameters for instantiating a new DBService.
type Config struct {
	MongoClient mongodb.Client
}

// New returns a new DBService client.
func New(ctx context.Context, config *Config) (*DBService, error) {
	return &DBService{Client: config.MongoClient}, nil
}

// Close closes the DBService connection with the client.
func (s *DBService) Close(ctx context.Context) {
	if err := s.Client.Disconnect(ctx); err != nil {
		logFatalf("failed closing connection with the DBService client: %s", err)
	}
}

// Test returns a test response.
func (s *DBService) AddUser(ctx context.Context, req *dpb.AddUserRequest) (*dpb.AddUserResponse, error) {
	err := s.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "failed to connect to MongoDB client: %s", err)
	}
	collection := s.Client.Database("db").Collection("users")
	res, err := collection.InsertOne(ctx, bson.M{"name": req.Name, "surname": req.Surname})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert to MongoDB: %s", err)
	}
	insertedID := res.InsertedID.(primitive.ObjectID)
	return &dpb.AddUserResponse{Id: insertedID.Hex()}, nil
}
