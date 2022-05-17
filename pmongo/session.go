package pmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SessionInterface interface {
	StartTransaction(...*options.TransactionOptions) error
	AbortTransaction(context.Context) error
	CommitTransaction(context.Context) error
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error),
		opts ...*options.TransactionOptions) (interface{}, error)
	EndSession(context.Context)

	ClusterTime() bson.Raw
	OperationTime() *primitive.Timestamp
	Client() ClientInterface
	ID() bson.Raw

	AdvanceClusterTime(bson.Raw) error
	AdvanceOperationTime(*primitive.Timestamp) error
}

var _ SessionInterface = (*session)(nil)

type session struct {
	sess mongo.Session
}

func (s session) StartTransaction(options ...*options.TransactionOptions) error {
	return s.sess.StartTransaction(options...)
}

func (s session) AbortTransaction(ctx context.Context) error {
	return s.sess.AbortTransaction(ctx)
}

func (s session) CommitTransaction(ctx context.Context) error {
	return s.sess.CommitTransaction(ctx)
}

func (s session) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	return s.sess.WithTransaction(ctx, fn)
}

func (s session) ClusterTime() bson.Raw {
	return s.sess.ClusterTime()
}

func (s session) OperationTime() *primitive.Timestamp {
	return s.sess.OperationTime()
}

func (s session) Client() ClientInterface {
	return &Client{s.sess.Client()}
}

func (s session) ID() bson.Raw {
	return s.sess.ID()
}

func (s session) EndSession(ctx context.Context) {
	s.sess.EndSession(ctx)
}

func (s session) AdvanceClusterTime(raw bson.Raw) error {
	return s.sess.AdvanceClusterTime(raw)
}

func (s session) AdvanceOperationTime(timestamp *primitive.Timestamp) error {
	return s.sess.AdvanceOperationTime(timestamp)
}
