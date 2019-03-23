package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client satisfies the mongo.Client interface.
type Client interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	Database(name string, opts ...*options.DatabaseOptions) Database
}

// Database satisfies the mongo.Database interface.
type Database interface {
	Client() Client
	Collection(name string, opts ...*options.CollectionOptions) Collection
}

// Collection satisfies the mongo.Collection interface.
type Collection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Database() Database
}

func New(ctx context.Context, mongoAddress string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoAddress))
	if err != nil {
		return nil, err
	}
	return client, nil
}
