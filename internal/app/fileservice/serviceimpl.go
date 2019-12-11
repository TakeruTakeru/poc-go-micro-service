package fileservice

import (
	"context"

	fservice "github.com/TakeruTakeru/poc-go-micro-service/api/fileservice"
)

type FileService struct{}

func (fs *FileService) GetGoogleDriveFileList(context.Context, *GoogleDriveFileListRequest) (*GoogleDriveFileListResponse, error) {
	return &fservice.GoogleDriveFileListResponse{}
}

func (fs *FileService) GetGoogleStrageFileList(context.Context, *GoogleStrageFileListRequest) (*GoogleStrageFileListResponse, error) {
	return &fservice.GoogleStrageFileListResponse{}
}

func NewFileService() *FileService {
	return &FileService{}
}
