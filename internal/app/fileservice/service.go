package fileservice

import (
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
)

type FileHandler interface {
	CreateDir(string) error
	Upload(byte) (int32, error)
	Download(string) (int32, error)
	GetList(string) ([]*models.FileModel, error)
	GetFileInfo(string) (models.FileModel, error)
}
