package service

import (
	"context"
	"protobuf-master/proto"
	"testing"
	"time"

	"golang-project/db/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	dpb "golang-project/db/proto"
)

type fakeMongo struct{}

func (f fakeMongo) Connect(ctx context.Context) error {
	return nil
}

func (f fakeMongo) Disconnect(ctx context.Context) error {
	return nil
}

func (f fakeMongo) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return nil
}

func (f fakeMongo) Database(name string, opts ...*options.DatabaseOptions) mongodb.Database {
	return &fakeDatabase{}
}

type fakeDatabase struct{}

func (f fakeDatabase) Client() mongodb.Client {
	return nil
}

func (f fakeDatabase) Collection(text string, opts ...*options.CollectionOptions) mongodb.Collection {
	return &fakeCollection{}
}

type fakeCollection struct{}

func (f fakeCollection) InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	var id primitive.ObjectID = [12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	return &mongo.InsertOneResult{
		InsertedID: id,
	}, nil
}

func TestAddUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fakeDBServer, err := New(ctx, &Config{MongoClient: fakeMongo{}})
	if err != nil {
		t.Errorf("failed to start a fake MongoDB client")
	}

	tests := []struct {
		name string
		want *dpb.AddUserResponse
	}{
		{
			name: "Success",
			want: &dpb.AddUserResponse{Id: "010203040506070809000100"},
		},
	}
	for _, test := range tests {
		req := &dpb.AddUserRequest{Name: "name", Surname: "surname"}
		got, err := fakeDBServer.AddUser(ctx, req)
		if err != nil {
			t.Errorf("AddUser(%v) got unexpected error", test.name)
		}
		if !proto.Equal(test.want, got) {
			t.Errorf("AddUser(%v)= got %v wanted %v", test.name, got, test.want)
		}
	}
}
