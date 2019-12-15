package gstorage

import (
	"cloud.google.com/go/storage"
	"context"
	"log"
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
		log.Printf("Failed to create bucket: %v\n", err)
		return err
	}
	return nil
}

func NewGoogleStorageClient(ctx context.Context, projectId string, client *storage.Client) *GoogleStorageClient {
	return &GoogleStorageClient{ctx: ctx, projectId: projectId, client: client}
}
