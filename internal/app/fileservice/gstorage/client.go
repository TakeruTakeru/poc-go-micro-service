package gstorage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
	"log"
	"path/filepath"
)

type GoogleStorageClient struct {
	ctx       context.Context
	projectId string
	client    *storage.Client
}

func (gsc *GoogleStorageClient) CreateDir(dir string) error {
	bucket := gsc.client.Bucket(dir)
	if gsc.ctx == nil {
		log.Fatalf("client: %v", gsc.ctx)
	}
	if err := bucket.Create(gsc.ctx, gsc.projectId, nil); err != nil {
		fmt.Printf("Failed to create bucket: %v\n", err)
		return err
	}
	return nil
}

func (gsc *GoogleStorageClient) DeleteDir(dir string) error {
	bucket := gsc.client.Bucket(dir)
	if gsc.ctx == nil {
		log.Fatalf("client: %v", gsc.ctx)
	}
	if err := bucket.Delete(gsc.ctx); err != nil {
		fmt.Printf("Failed to delete bucket: %v\n", err)
		return err
	}
	return nil
}

func (gsc *GoogleStorageClient) Upload(fm *models.FileModel) (size int, err error) {
	bucket := gsc.client.Bucket(filepath.Dir(fm.Model.GetPath()))
	obj := bucket.Object(fm.Model.GetName())
	w := obj.NewWriter(gsc.ctx)
	if size, err = fmt.Fprintf(w, string(fm.Data)); err != nil {
		fmt.Printf("Failed to write object: %v\n", err)
	}
	if err = w.Close(); err != nil {
		fmt.Printf("Failed to close object: %v\n", err)
	}
	return
}

func NewGoogleStorageClient(ctx context.Context, projectId string, client *storage.Client) *GoogleStorageClient {
	return &GoogleStorageClient{ctx: ctx, projectId: projectId, client: client}
}
