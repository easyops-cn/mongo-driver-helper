package helper

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	c   *mongo.Client
	err error
}

func (c *client) Err() error {
	return c.err
}

func (c *client) Database(name string, opts ...*options.DatabaseOptions) DatabaseHelper {
	return &database{
		db: c.c.Database(name, opts...),
	}
}

func (c *client) Connect(ctx context.Context) {
	if c.err == nil {
		c.err = c.c.Connect(ctx)
	}
}

func (c *client) Disconnect(ctx context.Context) {
	if c.err == nil {
		c.err = c.c.Disconnect(ctx)
	}
}

func NewClient(opts ...*options.ClientOptions) ClientHelper {
	c, err := mongo.NewClient(opts...)
	return &client{
		c:   c,
		err: err,
	}
}
