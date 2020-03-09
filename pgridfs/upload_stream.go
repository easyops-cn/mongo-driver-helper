package pgridfs

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type UploadStreamInterface interface {
	Abort() error
	Close() error
	SetWriteDeadline(t time.Time) error
	Write(p []byte) (int, error)
}

var _ UploadStreamInterface = (*gridfs.UploadStream)(nil)
