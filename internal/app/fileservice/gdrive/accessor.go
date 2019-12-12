package gdrive

import (
	"time"

	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
)

func GetFile() (f *models.FileModel, e error) {
	f, e = models.NewFile("FileName", 1024, "/takeru/private/", time.Now(), time.Now(), "takeru", "love this file")
	return
}

func GetFileList() (li []*models.FileModel, e error) {
	f, e := models.NewFile("FileName", 1024, "/takeru/private/", time.Now(), time.Now(), "takeru", "love this file")
	if e != nil {
		return nil, e
	}
	for _, _ = range []int{1, 2, 3} {
		li = append(li, f)
	}
	return
}
