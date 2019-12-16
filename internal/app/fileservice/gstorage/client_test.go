package gstorage

import (
	"context"
	"testing"
	"time"

	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
)

// 環境変数の設定
// GOOGLE_CLOUD_KEYFILE_JSON="$(< /Users/takeru/go/src/google-app/storage/Application-559d1d1cc0a1.json)"

var (
	testdir = "test-dir021900"
)

func TestCreateGoogleStorageClient_正常系(t *testing.T) {
	conn := createConn()
	_, err := conn.NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
}

func TestCreateGoogleStorageClient_異常系(t *testing.T) {
	ctx := context.Background()
	conn := NewGoogleStorageConnector(ctx, "ERROR", "sodium-chalice-256606")
	_, err := conn.NewClient()
	if err == nil {
		t.Errorf(unexpectedError(err.Error()))
	}
}

func TestDeleteDir_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	err = client.DeleteDir(testdir)
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
}

func TestCreateDir_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	err = client.CreateDir(testdir)
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
}

func TestUpload_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	fm, _ := models.NewFile("test-file", 0, testdir+"/"+"test-file", time.Now(), time.Now(), "takeru", "")
	fm.Data = []byte("test data.\nテストデータ\n")
	_, err = client.Upload(fm)
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
}

func createConn() *GoogleStorageConnector {
	ctx := context.Background()
	return NewGoogleStorageConnector(ctx, "GOOGLE_CLOUD_KEYFILE_JSON", "sodium-chalice-256606")
}

func unexpectedError(mes string) string {
	return "Unexpected error detected: " + mes + "\n"
}
