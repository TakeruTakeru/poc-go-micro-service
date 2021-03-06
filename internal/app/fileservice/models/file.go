package models

import (
	"fmt"
	fservice "github.com/TakeruTakeru/poc-go-micro-service/api/fileservice"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type FileModel struct {
	Model *fservice.File
}

func (fm *FileModel) String() string {
	return fmt.Sprintf("Name: %s, Size: %d, Path: %s", fm.Model.GetName(), fm.Model.GetSize(), fm.Model.GetPath())
}

func (fm *FileModel) GetPath() string {
	return fm.Model.GetPath()
}

func (fm *FileModel) GetName() string {
	return fm.Model.GetName()
}

func (fm *FileModel) GetData() []byte {
	return fm.Model.GetData()
}

func (fm *FileModel) GetFullPath() string {
	return fm.GetPath() + fm.GetName()
}

func (fm *FileModel) GetCreator() string {
	return fm.Model.GetCreator()
}

func NewFile(
	name string, size int32, data []byte, path string, lastModTime time.Time,
	createdTime time.Time, creator string, desc string,
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
		Data:           data,
		Path:           path,
		LastModifiedAt: lastModAt,
		CreatedAt:      createdAt,
		Creator:        creator,
		Desc:           desc,
	}

	return &FileModel{
		Model: base,
	}, nil
}
