package service

import (
	"golang-project/db/mongodb"

	"google.golang.org/grpc/codes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc/status"

	"context"
	"log"

	dpb "golang-project/db/proto"
)

var (
	// Fatalf var for testing purposes.
	logFatalf = log.Fatalf
	db        = "db"
	usersDB   = "users"
)

// singleResults will contain results from FindOne method calls.
var singleResult struct {
	Value float64
}

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

// AddUser returns a AddUser response.
func (s *DBService) AddUser(ctx context.Context, req *dpb.AddUserRequest) (*dpb.AddUserResponse, error) {
	if err := s.Client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "failed to connect to MongoDB client: %s", err)
	}
	collection := s.Client.Database(db).Collection(usersDB)
	res, err := collection.InsertOne(ctx, bson.M{"name": req.Name, "surname": req.Surname})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert to MongoDB: %s", err)
	}
	insertedID := res.InsertedID.(primitive.ObjectID)
	return &dpb.AddUserResponse{Id: insertedID.Hex()}, nil
}

// GetUser returns a GetUser response.
func (s *DBService) GetUser(ctx context.Context, req *dpb.GetUserRequest) (*dpb.GetUserResponse, error) {
	if err := s.Client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "failed to connect to MongoDB client: %s", err)
	}
	collection := s.Client.Database(db).Collection(usersDB)
	if err := collection.FindOne(ctx, bson.M{"name": req.Name, "surname": req.Surname}).Decode(&singleResult); err == nil {
		return &dpb.GetUserResponse{
			Name:    req.Name,
			Surname: req.Surname,
		}, nil
	}
	return &dpb.GetUserResponse{}, status.Errorf(codes.NotFound, "could not find user matching request: %v", req)
}

// ListTables returns a ListTables response.
func (s *DBService) ListTables(ctx context.Context, req *dpb.ListTablesRequest) (*dpb.ListTablesResponse, error) {
	if err := s.Client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "failed to connect to MongoDB client: %s", err)
	}
	if tables, err := s.Client.Database(db).ListCollectionNames(ctx, bsonx.Doc{}); err == nil {
		return &dpb.ListTablesResponse{Tables: tables}, err
	}
	return &dpb.ListTablesResponse{}, status.Errorf(codes.NotFound, "could not find tables")
}
