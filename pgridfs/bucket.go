package pgridfs

import (
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/easyops-cn/mongo-driver-helper/pmongo"
)

type BucketInterface interface {
	Delete(fileID interface{}) error
	DownloadToStream(fileID interface{}, stream io.Writer) (int64, error)
	DownloadToStreamByName(filename string, stream io.Writer, opts ...*options.NameOptions) (int64, error)
	Drop() error
	Find(filter interface{}, opts ...*options.GridFSFindOptions) (pmongo.CursorInterface, error)
	OpenDownloadStream(fileID interface{}) (DownloadStreamInterface, error)
	OpenDownloadStreamByName(filename string, opts ...*options.NameOptions) (DownloadStreamInterface, error)
	OpenUploadStream(filename string, opts ...*options.UploadOptions) (UploadStreamInterface, error)
	OpenUploadStreamWithID(fileID interface{}, filename string, opts ...*options.UploadOptions) (UploadStreamInterface, error)
	Rename(fileID interface{}, newFilename string) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	UploadFromStream(filename string, source io.Reader, opts ...*options.UploadOptions) (primitive.ObjectID, error)
	UploadFromStreamWithID(fileID interface{}, filename string, source io.Reader, opts ...*options.UploadOptions) error
}

type Bucket struct {
	b *gridfs.Bucket
}

func (b *Bucket) Delete(fileID interface{}) error {
	return b.b.Delete(fileID)
}

func (b *Bucket) DownloadToStream(fileID interface{}, stream io.Writer) (int64, error) {
	return b.b.DownloadToStream(fileID, stream)
}

func (b *Bucket) DownloadToStreamByName(filename string, stream io.Writer, opts ...*options.NameOptions) (int64, error) {
	return b.b.DownloadToStreamByName(filename, stream, opts...)
}

func (b *Bucket) Drop() error {
	return b.b.Drop()
}

func (b *Bucket) Find(filter interface{}, opts ...*options.GridFSFindOptions) (pmongo.CursorInterface, error) {
	return b.b.Find(filter, opts...)
}

func (b *Bucket) OpenDownloadStream(fileID interface{}) (DownloadStreamInterface, error) {
	return b.b.OpenDownloadStream(fileID)
}

func (b *Bucket) OpenDownloadStreamByName(filename string, opts ...*options.NameOptions) (DownloadStreamInterface, error) {
	return b.b.OpenDownloadStreamByName(filename, opts...)
}

func (b *Bucket) OpenUploadStream(filename string, opts ...*options.UploadOptions) (UploadStreamInterface, error) {
	return b.b.OpenUploadStream(filename, opts...)
}

func (b *Bucket) OpenUploadStreamWithID(fileID interface{}, filename string, opts ...*options.UploadOptions) (UploadStreamInterface, error) {
	return b.b.OpenUploadStreamWithID(fileID, filename, opts...)
}

func (b *Bucket) Rename(fileID interface{}, newFilename string) error {
	return b.b.Rename(fileID, newFilename)
}

func (b *Bucket) SetReadDeadline(t time.Time) error {
	return b.b.SetReadDeadline(t)
}

func (b *Bucket) SetWriteDeadline(t time.Time) error {
	return b.b.SetWriteDeadline(t)
}

func (b *Bucket) UploadFromStream(filename string, source io.Reader, opts ...*options.UploadOptions) (primitive.ObjectID, error) {
	return b.b.UploadFromStream(filename, source, opts...)
}

func (b *Bucket) UploadFromStreamWithID(fileID interface{}, filename string, source io.Reader, opts ...*options.UploadOptions) error {
	return b.b.UploadFromStreamWithID(fileID, filename, source, opts...)
}
