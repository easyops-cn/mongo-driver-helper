package pmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ClientInterface interface {
	Connect(ctx context.Context) error
	Database(name string, opts ...*options.DatabaseOptions) DatabaseInterface
	Disconnect(ctx context.Context) error
	ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error)
	ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error)
	NumberSessionsInProgress() int
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	StartSession(opts ...*options.SessionOptions) (mongo.Session, error)
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error
	Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
}

var _ ClientInterface = (*Client)(nil)

type Client struct {
	c *mongo.Client
}

func NewClient(c *mongo.Client) ClientInterface {
	return &Client{
		c: c,
	}
}

func (c *Client) Connect(ctx context.Context) error {
	return c.c.Connect(ctx)
}

func (c *Client) Database(name string, opts ...*options.DatabaseOptions) DatabaseInterface {
	db := c.c.Database(name, opts...)
	return &Database{db: db}
}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.c.Disconnect(ctx)
}

func (c *Client) ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error) {
	return c.c.ListDatabaseNames(ctx, filter, opts...)
}

func (c *Client) ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error) {
	return c.c.ListDatabases(ctx, filter, opts...)
}

func (c *Client) NumberSessionsInProgress() int {
	return c.c.NumberSessionsInProgress()
}

func (c *Client) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return c.c.Ping(ctx, rp)
}

func (c *Client) StartSession(opts ...*options.SessionOptions) (mongo.Session, error) {
	return c.c.StartSession(opts...)
}

func (c *Client) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return c.c.UseSession(ctx, fn)
}

func (c *Client) UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error {
	return c.c.UseSessionWithOptions(ctx, opts, fn)
}

func (c *Client) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return c.c.Watch(ctx, pipeline, opts...)
}
