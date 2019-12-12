package fileservice

import (
	"context"

	"github.com/TakeruTakeru/poc-go-micro-service/api/fileservice"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/gdrive"
)

type FileService struct{}

func (fs *FileService) GetGoogleDriveFileList(c context.Context, req *fileservice.GoogleDriveFileListRequest) (*fileservice.GoogleDriveFileListResponse, error) {
	var response []*fileservice.File
	files, e := gdrive.GetFileList()
	if e != nil {
		return nil, e
	}
	for _, file := range files {
		response = append(response, file.Model)
	}
	return &fileservice.GoogleDriveFileListResponse{File: response, RequestAt: nil}, nil
}

func (fs *FileService) GetGoogleStrageFileList(c context.Context, req *fileservice.GoogleStrageFileListRequest) (*fileservice.GoogleStrageFileListResponse, error) {
	var response []*fileservice.File
	files, e := gdrive.GetFileList()
	if e != nil {
		return nil, e
	}
	for _, file := range files {
		response = append(response, file.Model)
	}
	return &fileservice.GoogleStrageFileListResponse{File: response, RequestAt: nil}, nil
}

func NewFileService() *FileService {
	return &FileService{}
}
