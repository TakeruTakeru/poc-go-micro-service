package fileservice

import (
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
)

type FileHandler interface {
	CreateDir(string) error
	DeleteDir(string) error
	Upload(*models.FileModel) (int, error)
	Download(string) (int, error)
	GetList(string) ([]*models.FileModel, error)
	GetFileInfo(string) (models.FileModel, error)
}
