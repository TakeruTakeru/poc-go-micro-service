package gstorage

import (
	"context"
	"testing"
)

// 環境変数の設定
// GOOGLE_CLOUD_KEYFILE_JSON="$(< /Users/takeru/go/src/google-app/storage/Application-559d1d1cc0a1.json)"

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

func TestCreateDir(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	err = client.CreateDir("takeru-storage")
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
