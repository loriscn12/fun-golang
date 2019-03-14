package service

import (
<<<<<<< HEAD
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
		
=======
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	"go.mongodb.org/mongo-driver/mongo/options"
	
>>>>>>> b4dfd8f6f2b6bc2534d4c6cdbc88be08ea0cac2b
	"log"
	"context"

	dpb "golang-project/db/proto"
<<<<<<< HEAD
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
=======
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
>>>>>>> b4dfd8f6f2b6bc2534d4c6cdbc88be08ea0cac2b
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
<<<<<<< HEAD
=======
	grpcServer := grpc.NewServer()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed to instantiate a new MongoDB client: %s", err)
	}
	dpb.RegisterDatabaseServer(grpcServer, &dbServer{dbClient: client})
	grpcServer.Serve(lis)
>>>>>>> b4dfd8f6f2b6bc2534d4c6cdbc88be08ea0cac2b
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