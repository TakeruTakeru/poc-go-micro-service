package gstorage

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
	"google.golang.org/api/iterator"
)

type GoogleStorageClient struct {
	ctx       context.Context
	projectId string
	client    *storage.Client
}

//param dir is almost same as mkdir -p dir.
//for exmaple, param:dir is constructed with /<bucket_name>/<dir>/<dir>
func (gsc *GoogleStorageClient) CreateDir(dir string) (err error) {
	bname, fname := gsc.separateBucketNameAndFileName(dir)
	bucket := gsc.client.Bucket(bname)
	if fname != "" {
		obj := bucket.Object(fname + "/")
		w := obj.NewWriter(gsc.ctx)
		if err = w.Close(); err != nil {
			fmt.Printf("Failed to close object: %v\n", err)
			return
		}
		return
	}
	if err = bucket.Create(gsc.ctx, gsc.projectId, nil); err != nil {
		fmt.Printf("Failed to create bucket: %v\n", err)
		return
	}
	return
}

func (gsc *GoogleStorageClient) DeleteDir(dir string) error {
	bucket := gsc.client.Bucket(dir)
	if err := bucket.Delete(gsc.ctx); err != nil {
		fmt.Printf("Failed to delete bucket: %v\n", err)
		return err
	}
	return nil
}

func (gsc *GoogleStorageClient) Upload(fm *models.FileModel) (size int, err error) {
	bname, fname := gsc.separateBucketNameAndFileName(fm.Model.GetPath())
	bucket := gsc.client.Bucket(bname)
	objname := fm.Model.GetName()
	if fname != "" {
		objname = fname + "/" + objname
	}
	obj := bucket.Object(objname)
	w := obj.NewWriter(gsc.ctx)
	if size, err = fmt.Fprintf(w, string(fm.Data)); err != nil {
		fmt.Printf("Failed to write object: %v\n", err)
	}
	if err = w.Close(); err != nil {
		fmt.Printf("Failed to close object: %v\n", err)
	}
	return
}

func (gsc *GoogleStorageClient) Delete(path string) (err error) {
	bucket := gsc.client.Bucket(filepath.Dir(path))
	fname := filepath.Base(path)
	obj := bucket.Object(fname)
	err = obj.Delete(gsc.ctx)
	return
}

func (gsc *GoogleStorageClient) Download(path string) (fm *models.FileModel, err error) {
	fname := filepath.Base(path)
	obj := gsc.client.Bucket(filepath.Dir(path)).Object(fname)
	attr, err := obj.Attrs(gsc.ctx)
	if err != nil {
		return
	}
	rc, err := obj.NewReader(gsc.ctx)
	if err != nil {
		return
	}
	defer rc.Close()
	data, err := ioutil.ReadAll(rc)
	fm, _ = models.NewFile(attr.Name, int32(attr.Size), attr.Bucket, attr.Updated, attr.Created, attr.Owner, "")
	fm.Data = data
	return
}

func (gsc *GoogleStorageClient) GetBucketList(path string) (dirs []string, err error) {
	var battrs *storage.BucketAttrs
	var iterr error

	it := gsc.client.Buckets(gsc.ctx, gsc.projectId)
	for {
		battrs, iterr = it.Next()
		if iterr == iterator.Done {
			break
		}
		if iterr != nil {
			err = iterr
			return
		}
		dirs = append(dirs, battrs.Name)
	}
	return
}

func (gsc *GoogleStorageClient) GetFileList(path string) (files []*models.FileModel, err error) {
	var attrs *storage.ObjectAttrs
	var iterr error
	bucket := gsc.client.Bucket(path)
	it := bucket.Objects(gsc.ctx, nil)
	for {
		attrs, iterr = it.Next()
		if iterr == iterator.Done {
			break
		}
		if iterr != nil {
			err = iterr
			return
		}
		fm, _ := models.NewFile(attrs.Name, int32(attrs.Size), attrs.Bucket, attrs.Updated, attrs.Created, attrs.Owner, "")
		files = append(files, fm)
	}
	return
}

func (gsc *GoogleStorageClient) GetFileInfo(path string) (fm *models.FileModel, err error) {
	bucket := gsc.client.Bucket(filepath.Dir(path))
	obj := bucket.Object(filepath.Base(path))
	attr, err := obj.Attrs(gsc.ctx)
	if err != nil {
		return
	}
	fm, err = models.NewFile(attr.Name, int32(attr.Size), path, attr.Updated, attr.Created, attr.Owner, "")
	return
}

func (gsc *GoogleStorageClient) separateBucketNameAndFileName(path string) (bname string, fname string) {
	elem := strings.Split(path, "/")
	if len(elem) > 1 {
		lastIdx := len(elem) - 1
		if elem[lastIdx] == "" {
			lastIdx = lastIdx - 1
		}
		bname = elem[0]
		fname = strings.Join(elem[1:lastIdx+1], "/")
		return
	} else {
		bname = path
		fname = ""
	}
	return
}

func NewGoogleStorageClient(ctx context.Context, projectId string, client *storage.Client) *GoogleStorageClient {
	return &GoogleStorageClient{ctx: ctx, projectId: projectId, client: client}
}
