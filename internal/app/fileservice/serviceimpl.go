package fileservice

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/TakeruTakeru/poc-go-micro-service/api/fileservice"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/gdrive"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/gstorage"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
	"github.com/TakeruTakeru/poc-go-micro-service/pkg/logger"
	"github.com/golang/protobuf/ptypes"
)

type FileService struct{}

var (
	PROJECT_ID_ENV    = "PROJECT_ID"
	KEY_FILE_PATH_ENV = "GOOGLE_CLOUD_KEYFILE_JSON"
)

func (fs *FileService) GetGoogleDriveFileList(c context.Context, req *fileservice.GoogleDriveFileListRequest) (*fileservice.GoogleDriveFileListResponse, error) {
	files, e := gdrive.GetFileList()
	if e != nil {
		return nil, e
	}
	return &fileservice.GoogleDriveFileListResponse{Path: files, RequestAt: ptypes.TimestampNow()}, nil
}

func (fs *FileService) GetGoogleStorageFileList(c context.Context, req *fileservice.GoogleStorageFileListRequest) (res *fileservice.GoogleStorageFileListResponse, err error) {
	conn := gstorage.NewGoogleStorageConnector(context.Background(), KEY_FILE_PATH_ENV, PROJECT_ID_ENV)
	client, err := conn.NewClient()
	if err != nil {
		return
	}
	files, err := client.GetFileList(req.GetPath())
	return &fileservice.GoogleStorageFileListResponse{Path: files, RequestAt: ptypes.TimestampNow()}, nil
}

func (fs *FileService) GetGoogleStorageFile(ctx context.Context, req *fileservice.GoogleStorageFileRequest) (res *fileservice.GoogleStorageFileResponse, err error) {
	conn := gstorage.NewGoogleStorageConnector(context.Background(), KEY_FILE_PATH_ENV, PROJECT_ID_ENV)
	client, err := conn.NewClient()
	if err != nil {
		return
	}
	// grpc-gatewayだと要件を満たすようなURLマッピングができなかったので、google-storageのパスは
	// base64エンコーディング(URL safe)されたものが送られてくるように実装
	base64Id := req.GetBase64Id()
	decodedId, err := base64.URLEncoding.DecodeString(base64Id)
	fm, err := client.Download(string(decodedId))
	if err != nil {
		return
	}
	res = &fileservice.GoogleStorageFileResponse{File: fm.Model}
	return
}

func (fs *FileService) UploadGoogleStorageFile(ctx context.Context, req *fileservice.GoogleStorageFileUploadRequest) (res *fileservice.GoogleStorageFileUploadResponse, err error) {
	conn := gstorage.NewGoogleStorageConnector(context.Background(), KEY_FILE_PATH_ENV, PROJECT_ID_ENV)
	client, err := conn.NewClient()
	if err != nil {
		return
	}
	cert := req.GetCert()
	ok := verify(cert)
	if !ok {
		return nil, fmt.Errorf("Invalid Access.")
	}
	fileReq := req.GetFile()
	lastModAt, err := ptypes.Timestamp(fileReq.GetLastModifiedAt())
	createdAt, err := ptypes.Timestamp(fileReq.GetCreatedAt())
	fm, err := models.NewFile(
		fileReq.GetName(),
		fileReq.GetSize(),
		fileReq.GetData(),
		fileReq.GetPath(),
		lastModAt,
		createdAt,
		fileReq.GetCreator(),
		fileReq.GetDesc(),
	)
	if err != nil {
		return
	}
	_, err = client.Upload(fm)
	if err != nil {
		return
	}
	return
}

func NewFileService() *FileService {
	return &FileService{}
}

func verify(cert *fileservice.Certification) bool {
	user := cert.GetUser()
	pass := cert.GetPassword()
	if user != os.Getenv("USER") {
		logger.Errorf("Invalid User Detected")
		return false
	}
	if pass != os.Getenv("PASSWORD") {
		logger.Errorf("Invalid Password Detected")
		return false
	}
	return true
}
