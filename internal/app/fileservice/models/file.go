package models

import (
	fservice "github.com/TakeruTakeru/poc-go-micro-service/api/fileservice"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type FileModel struct {
	Model *fservice.File
	Data  []byte
}

func NewFile(
	name string, size int32, path string, lastModTime time.Time,
	createdTime time.Time, createor string, desc string,
) (*FileModel, error) {
	lastModAt, err := ptypes.TimestampProto(lastModTime)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.TimestampProto(createdTime)
	if err != nil {
		return nil, err
	}
	base := &fservice.File{
		Name:           name,
		Size:           size,
		Path:           path,
		LastModifiedAt: lastModAt,
		CreatedAt:      createdAt,
		Creator:        createor,
		Desc:           desc,
	}

	return &FileModel{
		Model: base,
	}, nil
}
