package service

import (
	"google.golang.org/grpc"
	"log"
	"time"
	"testing"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc/test/grpc_testing"
)

type FakeMongo struct{}

func (f FakeMongo) Connect(ctx context.Context) error{
	return nil
}

func (f FakeMongo) Disconnect(ctx context.Context) error{
	return nil
}

func (f FakeMongo) Ping(ctx context.Context, rp *readpref.ReadPref) error{
	return nil
}

func (f FakeMongo) StartSession(opts ...*options.SessionOptions) (mongo.Session, error){
	return nil, nil
}

func (f FakeMongo) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database{
	return nil
}

func (f FakeMongo) ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error){
	return mongo.ListDatabasesResult{}, nil
}

func (f FakeMongo) ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error){
	return []string{}, nil
}

func (f FakeMongo) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error{
	return nil
}

func (f FakeMongo) UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error{
	return nil
}

func (f FakeMongo) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error){
	return &mongo.ChangeStream{}, nil
}


func TestAddUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel();
	//lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 10001))

	testMongo := &FakeMongo{}
	fakeClient := testMongo.Connect(ctx)
	dbServer, err := New(ctx, &Config{MongoClient: testMongo})
	if err != nil {
		log.Fatalf("failed to instanciate a new DBService client: %s", err)
	}
	defer dbServer.Close(ctx)
	grpcServer := grpc.NewServer()
	
	grpc_testing.RegisterTestServiceServer(grpcServer, &DBService{})
	//grpc_testing.Serve(lis)


	//TODO(loris): Add test for AddUser function.
}