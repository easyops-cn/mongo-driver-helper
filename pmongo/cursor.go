package pmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type CursorInterface interface {
	All(ctx context.Context, results interface{}) error
	Close(ctx context.Context) error
	Decode(val interface{}) error
	Err() error
	ID() int64
	Next(ctx context.Context) bool
	TryNext(ctx context.Context) bool
}

var _ CursorInterface = (*mongo.Cursor)(nil)
