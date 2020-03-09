package helper

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type database struct {
	db *mongo.Database
}

func (d *database) Collection(name string, opts ...*options.CollectionOptions) CollectionHelper {
	return &collection{
		coll: d.db.Collection(name, opts...),
	}
}
