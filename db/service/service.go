package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
		
	"log"
	"context"

	dpb "golang-project/db/proto"
)

type DBService struct {
	Client *mongo.Client
}

type Config struct {
	MongoAddress string
}

// New returns a new DBService client.
func New(ctx context.Context, config *Config)(*DBService, error){
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoAddress))
	if err != nil {
		return &DBService{}, err
	}
	return &DBService{Client: client}, nil
}

// Close closes the DBService connection with the client.
func (s *DBService) Close(ctx context.Context) {
	if err := s.Client.Disconnect(ctx); err != nil {
		log.Fatalf("failed closing connection with the DBService client: %s", err)
	}
}

// Test returns a test response.
func (s *DBService) AddUser(ctx context.Context, req *dpb.AddUserRequest) (*dpb.AddUserResponse, error) {
	err := s.Client.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatalf("failed to connect to MongoDB client: %s", err)
	}
	collection := s.Client.Database("db").Collection("users")
	res, err := collection.InsertOne(ctx, bson.M{"name": req.Name, "surname": req.Surname})
	if err != nil{
		log.Fatalf("failed to insert to MongoDB: %s", err)
	}
	insertedID := res.InsertedID.(primitive.ObjectID)
	return &dpb.AddUserResponse{Id: insertedID.Hex()}, nil
}