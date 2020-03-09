package pgridfs

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type DownloadStreamInterface interface {
	Close() error
	Read(p []byte) (int, error)
	SetReadDeadline(t time.Time) error
	Skip(skip int64) (int64, error)
}

var _ DownloadStreamInterface = (*gridfs.DownloadStream)(nil)
