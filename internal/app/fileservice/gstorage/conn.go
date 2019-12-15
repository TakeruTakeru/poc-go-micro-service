package gstorage

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GoogleStorageConnector struct {
	ctx       context.Context
	jsonPath  string
	projectId string
}

func (gsc *GoogleStorageConnector) NewClient() (client *GoogleStorageClient, err error) {
	googleClient, err := storage.NewClient(gsc.ctx, option.WithCredentialsJSON([]byte(os.Getenv(gsc.jsonPath))))
	if err != nil {
		log.Printf("Failed to create client: %v\n", err)
		return nil, err
	}
	if gsc.ctx == nil {
		log.Fatalf("conn: %v", gsc.ctx)
	}
	client = NewGoogleStorageClient(gsc.ctx, gsc.projectId, googleClient)
	return
}

func NewGoogleStorageConnector(ctx context.Context, jsonPath string, projectId string) (conn *GoogleStorageConnector) {
	conn = &GoogleStorageConnector{
		ctx:       ctx,
		jsonPath:  jsonPath,
		projectId: projectId,
	}
	return
}
