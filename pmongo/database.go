package pmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type DatabaseInterface interface {
	Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (CursorInterface, error)
	Client() ClientInterface
	Collection(name string, opts ...*options.CollectionOptions) CollectionInterface
	Drop(ctx context.Context) error
	ListCollectionNames(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) ([]string, error)
	ListCollections(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) (CursorInterface, error)
	Name() string
	ReadConcern() *readconcern.ReadConcern
	ReadPreference() *readpref.ReadPref
	RunCommand(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) SingleResultInterface
	RunCommandCursor(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) (CursorInterface, error)
	Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
	WriteConcern() *writeconcern.WriteConcern

	NewBucket(opts ...*options.BucketOptions) (BucketInterface, error)
}

var _ DatabaseInterface = (*Database)(nil)

type Database struct {
	db *mongo.Database
}

func (d *Database) NewBucket(opts ...*options.BucketOptions) (BucketInterface, error) {
	buck, err := gridfs.NewBucket(d.db, opts...)
	if err != nil {
		return nil, err
	}
	return &Bucket{buck}, nil
}

func (d *Database) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (CursorInterface, error) {
	return d.db.Aggregate(ctx, pipeline, opts...)
}

func (d *Database) Client() ClientInterface {
	c := d.db.Client()
	return &Client{c: c}
}

func (d *Database) Collection(name string, opts ...*options.CollectionOptions) CollectionInterface {
	coll := d.db.Collection(name, opts...)
	return &Collection{coll: coll}
}

func (d *Database) Drop(ctx context.Context) error {
	return d.db.Drop(ctx)
}

func (d *Database) ListCollectionNames(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) ([]string, error) {
	return d.db.ListCollectionNames(ctx, filter, opts...)
}

func (d *Database) ListCollections(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) (CursorInterface, error) {
	return d.db.ListCollections(ctx, filter, opts...)
}

func (d *Database) Name() string {
	return d.db.Name()
}

func (d *Database) ReadConcern() *readconcern.ReadConcern {
	return d.db.ReadConcern()
}

func (d *Database) ReadPreference() *readpref.ReadPref {
	return d.db.ReadPreference()
}

func (d *Database) RunCommand(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) SingleResultInterface {
	return d.db.RunCommand(ctx, runCommand, opts...)
}

func (d *Database) RunCommandCursor(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) (CursorInterface, error) {
	return d.db.RunCommandCursor(ctx, runCommand, opts...)
}

func (d *Database) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return d.db.Watch(ctx, pipeline, opts...)
}

func (d *Database) WriteConcern() *writeconcern.WriteConcern {
	return d.db.WriteConcern()
}
