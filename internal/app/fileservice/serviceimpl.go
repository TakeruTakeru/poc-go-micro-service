package fileservice

import (
	"context"
	"time"

	"github.com/TakeruTakeru/poc-go-micro-service/api/fileservice"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/gdrive"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/gstorage"
	"github.com/golang/protobuf/ptypes"
)

type FileService struct{}

var (
	PROJECT_ID        = "sodium-chalice-256606"
	KEY_FILE_PATH_ENV = "GOOGLE_CLOUD_KEYFILE_JSON"
)

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

func (fs *FileService) GetGoogleStrageFileList(c context.Context, req *fileservice.GoogleStrageFileListRequest) (res *fileservice.GoogleStrageFileListResponse, err error) {
	conn := gstorage.NewGoogleStorageConnector(context.Background(), KEY_FILE_PATH_ENV, PROJECT_ID)
	client, err := conn.NewClient()
	if err != nil {
		return
	}
	var files []*fileservice.File
	fms, err := client.GetFileList(req.GetPath())
	for _, fm := range fms {
		files = append(files, fm.Model)
	}
	requestedAt, _ := ptypes.TimestampProto(time.Now())
	return &fileservice.GoogleStrageFileListResponse{File: files, RequestAt: requestedAt}, nil
}

func NewFileService() *FileService {
	return &FileService{}
}
