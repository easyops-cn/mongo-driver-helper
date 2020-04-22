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
	indexView mongo.IndexView
}

func NewIndexView(indexView mongo.IndexView) IndexView {
	return IndexView{indexView: indexView}
}

func (i IndexView) List(ctx context.Context, opts ...*options.ListIndexesOptions) (CursorInterface, error) {
	return i.indexView.List(ctx, opts...)
}

func (i IndexView) CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return i.indexView.CreateOne(ctx, model, opts...)
}

func (i IndexView) CreateMany(ctx context.Context, models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	return i.indexView.CreateMany(ctx, models, opts...)
}

func (i IndexView) DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return i.indexView.DropOne(ctx, name, opts...)
}

func (i IndexView) DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return i.indexView.DropAll(ctx, opts...)
}

var _ IndexViewInterface = IndexView{}
