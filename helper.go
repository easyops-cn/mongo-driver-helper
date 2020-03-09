package helper

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientHelper interface {
	Database(name string, opts ...*options.DatabaseOptions) DatabaseHelper
	Connect(ctx context.Context)
	Disconnect(ctx context.Context)
	Err() error
}

type DatabaseHelper interface {
	Collection(name string, opts ...*options.CollectionOptions) CollectionHelper
}

type CollectionHelper interface {
	BulkWrite(ctx context.Context, models []mongo.WriteModel,
		opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, documents []interface{},
		opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (CursorHelper, error)
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) SingleResultHelper
}

type CursorHelper interface {
	ID() int64
	Next(ctx context.Context) bool
	TryNext(ctx context.Context) bool
	Decode(val interface{}) error
	Err() error
	Close(ctx context.Context) error
	All(ctx context.Context, results interface{}) error
}

type SingleResultHelper interface {
	Decode(v interface{}) error
	DecodeBytes() (bson.Raw, error)
	Err() error
}
