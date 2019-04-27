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
	pingErr       error
	insertErr     error
	disconnectErr error
	findErr       error
}

func (f fakeMongo) Connect(ctx context.Context) error {
	return nil
}

func (f fakeMongo) Disconnect(ctx context.Context) error {
	return f.disconnectErr
}

func (f fakeMongo) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return f.pingErr
}

func (f fakeMongo) Database(name string, opts ...*options.DatabaseOptions) mongodb.Database {
	return &fakeDatabase{client: f}
}

type fakeDatabase struct {
	client fakeMongo
}

func (f fakeDatabase) Client() mongodb.Client {
	return f.client
}

func (f fakeDatabase) Collection(text string, opts ...*options.CollectionOptions) mongodb.Collection {
	return &fakeCollection{client: f.client}
}

type fakeCollection struct {
	client fakeMongo
}

func (f fakeCollection) InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	var id primitive.ObjectID = [12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	if f.client.insertErr != nil {
		return nil, f.client.insertErr
	}
	return &mongo.InsertOneResult{
		InsertedID: id,
	}, f.client.insertErr
}

func (f fakeCollection) FindOne(context.Context, interface{}, ...*options.FindOneOptions) mongodb.SingleResult {
	return &fakeSingleResult{err: f.client.findErr}
}

func (f fakeCollection) Database() mongodb.Database {
	return &fakeDatabase{}
}

type fakeSingleResult struct {
	mongodb.SingleResult
	err error
}

func (f fakeSingleResult) Decode(v interface{}) error {
	return f.err
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

	errs := []string{}
	logFatalf = func(format string, args ...interface{}) {
		if len(args) > 0 {
			errs = append(errs, fmt.Sprintf(format, args))
		} else {
			errs = append(errs, format)
		}
	}
	tests := []struct {
		name          string
		disconnectErr error
		returnErr     bool
		wantCode      codes.Code
	}{
		{
			name: "Success",
		},
		{
			name:          "Fatal failure due to Disconnect error",
			returnErr:     true,
			disconnectErr: errors.New("returns error"),
		},
	}
	for _, test := range tests {
		mockedMongo.disconnectErr = test.disconnectErr
		fakeDBServer.Close(ctx)
		if test.returnErr && (len(errs) != 1) {
			t.Errorf("Close(%s) expected one error, got %d", test.name, len(errs))
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
		pingErr   error
		insertErr error
		returnErr bool
		wantCode  codes.Code
		wantErr   string
	}{
		{
			name: "Success",
			want: &dpb.AddUserResponse{Id: "010203040506070809000100"},
		},
		{
			name:      "Failure due to Ping error",
			returnErr: true,
			pingErr:   errors.New("returns error"),
			wantCode:  codes.FailedPrecondition,
			wantErr:   "rpc error: code = FailedPrecondition desc = failed to connect to MongoDB client: returns error",
		},
		{
			name:      "Failure due to InsertOne error",
			returnErr: true,
			insertErr: errors.New("returns error"),
			wantCode:  codes.Internal,
			wantErr:   "rpc error: code = Internal desc = failed to insert to MongoDB: returns error",
		},
	}
	for _, test := range tests {
		mockedMongo.pingErr = test.pingErr
		mockedMongo.insertErr = test.insertErr
		req := &dpb.AddUserRequest{Name: "name", Surname: "surname"}
		got, err := fakeDBServer.AddUser(ctx, req)
		gotStatus, _ := status.FromError(err)
		gotCode := gotStatus.Code()
		if test.returnErr {
			if gotCode != test.wantCode || err.Error() != test.wantErr {
				t.Errorf("AddUser(%v) got status: %v error: %v want status: %v error: %v", test.name, gotCode, err.Error(), test.wantCode, test.wantErr)
			}
		} else if !proto.Equal(test.want, got) {
			t.Errorf("AddUser(%v)= got %v wanted %v", test.name, got, test.want)
		}
	}
}

func TestGetUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mockedMongo := &fakeMongo{}
	fakeDBServer, err := New(ctx, &Config{MongoClient: mockedMongo})
	if err != nil {
		t.Fatalf("failed to start a fake MongoDB client")
	}
	req := &dpb.GetUserRequest{Name: "name", Surname: "surname"}
	tests := []struct {
		name      string
		want      *dpb.GetUserResponse
		pingErr   error
		findErr   error
		returnErr bool
		wantCode  codes.Code
		wantErr   string
	}{
		{
			name: "Success",
			want: &dpb.GetUserResponse{
				Name:    "name",
				Surname: "surname",
			},
		},
		{
			name:      "Failure due to Ping error",
			returnErr: true,
			pingErr:   errors.New("returns error"),
			wantCode:  codes.FailedPrecondition,
			wantErr:   "rpc error: code = FailedPrecondition desc = failed to connect to MongoDB client: returns error",
		},
		{
			name:      "Failure due to FindOne error",
			returnErr: true,
			findErr:   errors.New("returns error"),
			wantCode:  codes.NotFound,
			wantErr:   fmt.Sprintf("rpc error: code = NotFound desc = could not find user matching request: %v", req),
		},
	}
	for _, test := range tests {
		mockedMongo.pingErr = test.pingErr
		mockedMongo.findErr = test.findErr
		got, err := fakeDBServer.GetUser(ctx, req)
		gotStatus, _ := status.FromError(err)
		gotCode := gotStatus.Code()
		if test.returnErr {
			if gotCode != test.wantCode || err.Error() != test.wantErr {
				t.Errorf("GetUser(%v) got status: %v error: %v want status: %v error: %v", test.name, gotCode, err.Error(), test.wantCode, test.wantErr)
			}
		} else if !proto.Equal(test.want, got) {
			t.Errorf("GetUser(%v) got %v wanted %v", test.name, got, test.want)
		}
	}
}
