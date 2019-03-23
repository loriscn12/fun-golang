package service

import (
	"context"
	"errors"
	"fmt"
	"protobuf-master/proto"
	"testing"
	"time"

	"golang-project/db/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	dpb "golang-project/db/proto"
)

type fakeMongo struct {
	mongodb.Client
	client        mongodb.Client
	pingErr       bool
	insertErr     bool
	disconnectErr bool
}

func (f fakeMongo) Connect(ctx context.Context) error {
	return nil
}

func (f fakeMongo) Disconnect(ctx context.Context) error {
	if f.disconnectErr {
		return errors.New("returns error")
	}
	return nil
}

func (f fakeMongo) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	if f.pingErr {
		return errors.New("returns error")
	}
	return nil
}

func (f fakeMongo) Database(name string, opts ...*options.DatabaseOptions) mongodb.Database {
	return &fakeDatabase{
		client: f,
	}
}

type fakeDatabase struct {
	mongodb.Database
	client fakeMongo
}

func (f fakeDatabase) Client() mongodb.Client {
	return &fakeMongo{
		client: f.client,
	}
}

func (f fakeDatabase) Collection(text string, opts ...*options.CollectionOptions) mongodb.Collection {
	return &fakeCollection{
		client: f.client,
	}
}

type fakeCollection struct {
	mongodb.Collection
	client fakeMongo
}

func (f fakeCollection) InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	var id primitive.ObjectID = [12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	if f.client.insertErr {
		return nil, errors.New("returns error")
	}
	return &mongo.InsertOneResult{
		InsertedID: id,
	}, nil
}

func (f fakeCollection) Database() mongodb.Database {
	return &fakeDatabase{}
}

func TestClose(t *testing.T) {
	realFatalf := logFatalf
	defer func() {
		logFatalf = realFatalf
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mockedMongo := &fakeMongo{}
	fakeDBServer, err := New(ctx, &Config{MongoClient: mockedMongo})
	if err != nil {
		t.Fatalf("failed to start a fake MongoDB client")
	}

	errors := []string{}
	logFatalf = func(format string, args ...interface{}) {
		if len(args) > 0 {
			errors = append(errors, fmt.Sprintf(format, args))
		} else {
			errors = append(errors, format)
		}
	}
	tests := []struct {
		name          string
		disconnectErr bool
		returnErr     bool
		wantCode      codes.Code
	}{
		{
			name: "Success",
		},
		{
			name:          "Fatal failure due to Disconnect error",
			returnErr:     true,
			disconnectErr: true,
		},
	}
	for _, test := range tests {
		mockedMongo.disconnectErr = test.disconnectErr
		fakeDBServer.Close(ctx)
		if test.returnErr && (len(errors) != 1) {
			t.Errorf("Close(%s) expected one error, got %d", test.name, len(errors))
		}
	}
}

func TestAddUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mockedMongo := &fakeMongo{}
	fakeDBServer, err := New(ctx, &Config{MongoClient: mockedMongo})
	if err != nil {
		t.Fatalf("failed to start a fake MongoDB client")
	}

	tests := []struct {
		name      string
		want      *dpb.AddUserResponse
		pingErr   bool
		insertErr bool
		returnErr bool
		wantCode  codes.Code
	}{
		{
			name: "Success",
			want: &dpb.AddUserResponse{Id: "010203040506070809000100"},
		},
		{
			name:      "Failure due to Ping error",
			returnErr: true,
			pingErr:   true,
			wantCode:  codes.FailedPrecondition,
		},
		{
			name:      "Failure due to InsertOne error",
			returnErr: true,
			insertErr: true,
			wantCode:  codes.Internal,
		},
	}
	for _, test := range tests {
		mockedMongo.pingErr = test.pingErr
		mockedMongo.insertErr = test.insertErr
		req := &dpb.AddUserRequest{Name: "name", Surname: "surname"}
		got, err := fakeDBServer.AddUser(ctx, req)
		gotStatus, _ := status.FromError(err)
		gotCode := gotStatus.Code()
		if test.returnErr && (err != nil) == (gotCode != test.wantCode) {
			t.Errorf("AddUser(%v) got status %v want status %v", test.name, gotStatus, test.wantCode)
		} else if !proto.Equal(test.want, got) {
			t.Errorf("AddUser(%v)= got %v wanted %v", test.name, got, test.want)
		}
	}
}
