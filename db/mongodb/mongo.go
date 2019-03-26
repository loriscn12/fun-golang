package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	*mongo.Client
}

// Client satisfies the mongo.Client interface.
type Client interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	Database(name string, opts ...*options.DatabaseOptions) Database
}

func (c MongoClient) Database(name string, opts ...*options.DatabaseOptions) Database {
	return &MongoDatabase{Database: c.Client.Database(name, opts...)}
}

type MongoDatabase struct {
	*mongo.Database
}

// Database satisfies the mongo.Database interface.
type Database interface {
	Collection(name string, opts ...*options.CollectionOptions) Collection
}

func (d MongoDatabase) Collection(name string, opts ...*options.CollectionOptions) Collection {
	return &MongoCollection{Collection: d.Database.Collection(name, opts...)}
}

type MongoCollection struct {
	*mongo.Collection
}

// Collection satisfies the mongo.Collection interface.
type Collection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Database() Database
}

func (c MongoCollection) Database() Database {
	return &MongoDatabase{Database: c.Collection.Database()}
}

func New(ctx context.Context, mongoAddress string) (Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoAddress))
	if err != nil {
		return nil, err
	}
	return MongoClient{client}, nil
}
