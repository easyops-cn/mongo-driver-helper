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

func (s session) EndSession(ctx context.Context) {
	s.sess.EndSession(ctx)
}

func (s session) AdvanceClusterTime(raw bson.Raw) error {
	return s.sess.AdvanceClusterTime(raw)
}

func (s session) AdvanceOperationTime(timestamp *primitive.Timestamp) error {
	return s.sess.AdvanceOperationTime(timestamp)
}

func (s session) Client() ClientInterface {
	return &Client{s.sess.Client()}
}
