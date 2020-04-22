package pmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IndexViewInterface interface {
	List(ctx context.Context, opts ...*options.ListIndexesOptions) (CursorInterface, error)
	CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error)
	CreateMany(ctx context.Context, models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error)
	DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error)
	DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error)
}

type IndexView struct {
	mongo.IndexView
}

func (i IndexView) List(ctx context.Context, opts ...*options.ListIndexesOptions) (CursorInterface, error) {
	return i.IndexView.List(ctx, opts...)
}

var _ IndexViewInterface = IndexView{}
