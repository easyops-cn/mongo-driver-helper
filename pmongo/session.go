package pmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SessionInterface interface {
	StartTransaction(options ...*options.TransactionOptions) error
	AbortTransaction(ctx context.Context) error
	CommitTransaction(ctx context.Context) error
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error)
	ClusterTime() bson.Raw
	OperationTime() *primitive.Timestamp
	Client() ClientInterface
	EndSession(context.Context)

	AdvanceClusterTime(bson.Raw) error
	AdvanceOperationTime(*primitive.Timestamp) error
}

var _ SessionInterface = (*session)(nil)

type session struct {
	mongo.Session
}

func (s session) Client() ClientInterface {
	return &Client{s.Session.Client()}
}
