package gstorage

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
)

// 環境変数の設定
// GOOGLE_CLOUD_KEYFILE_JSON="$(< /Users/takeru/go/src/google-app/storage/Application-559d1d1cc0a1.json)"

var (
	tempBucket   = "takeru-test-tempdir"
	parmBucket   = "takeru-test-parmanent"
	tempObj1     = "temp.txt"
	tempObj1Body = "test data.\nテストデータ\n"
	tempObj1Size int
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

func TestCreateDir_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	err = client.CreateDir(tempBucket)
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
}

func TestUpload_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	fm, _ := models.NewFile(tempObj1, 0, tempBucket+"/"+tempObj1, time.Now(), time.Now(), "takeru", "")
	fm.Data = []byte(tempObj1Body)
	tempObj1Size, err = client.Upload(fm)
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	fmt.Printf("Upload temp file. size: %d\n", tempObj1Size)
}

func TestDownload_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	fm, err := client.Download(tempBucket + "/" + tempObj1)
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	if string(fm.Data) != tempObj1Body {
		t.Errorf(unexpectedError(string(fm.Data)))
	}

	fmt.Printf("Download temp file. data:\n%s", string(fm.Data))
}

func TestGetDirList_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	dirs, err := client.GetDirList("")
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	fmt.Printf("Get directories list: %v\n", dirs)
}

func TestGetFileList_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	fms, err := client.GetFileList(tempBucket)
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	fmt.Printf("Get files list: %v\n", fms)
}

func TestGetFileInfo_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	fm, err := client.GetFileInfo(tempBucket + "/" + tempObj1)
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	fmt.Printf("Get file detail info: %v", fm)
}

func TestDelete_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	err = client.Delete(tempBucket + "/" + tempObj1)
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
}

func TestDeleteDir_正常系(t *testing.T) {
	client, err := createConn().NewClient()
	if err != nil {
		t.Errorf(unexpectedError(err.Error()))
	}
	err = client.DeleteDir(tempBucket)
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
