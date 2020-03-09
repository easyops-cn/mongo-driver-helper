package helper

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type collection struct {
	coll *mongo.Collection
}

func (c *collection) BulkWrite(ctx context.Context, models []mongo.WriteModel,
	opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return c.coll.BulkWrite(ctx, models, opts...)
}

func (c *collection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.coll.InsertOne(ctx, document, opts...)
}

func (c collection) InsertMany(ctx context.Context, documents []interface{},
	opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return c.coll.InsertMany(ctx, documents, opts...)
}

func (c collection) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (CursorHelper, error) {
	return c.coll.Find(ctx, filter, opts...)
}

func (c collection) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) SingleResultHelper {
	return c.coll.FindOne(ctx, filter, opts...)
}
