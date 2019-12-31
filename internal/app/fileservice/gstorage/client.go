package gstorage

import (
	"context"
	"fmt"
	"io/ioutil"

	"strings"

	"cloud.google.com/go/storage"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
	logger "github.com/TakeruTakeru/poc-go-micro-service/pkg/logger"
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
			logger.Errorf("Failed to close object: %v", err)
			return
		}
		logger.Printf("Created Directory '%s'", fname)
		return
	}
	if err = bucket.Create(gsc.ctx, gsc.projectId, nil); err != nil {
		logger.Errorf("Failed to create bucket: %v", err)
		return
	}
	logger.Printf("Created bucket '%s'", bname)
	return
}

func (gsc *GoogleStorageClient) DeleteBucket(bname string) (err error) {
	bucket := gsc.client.Bucket(bname)
	paths, err := gsc.GetFileList(bname)
	if len(paths) > 1 {
		for _, path := range paths {
			err = gsc.Delete(path)
		}
		if err = bucket.Delete(gsc.ctx); err != nil {
			logger.Errorf("Failed to delete bucket: %v", err)
			return
		}
	}
	logger.Printf("Delete bucket '%s'", bname)
	return
}

func (gsc *GoogleStorageClient) Upload(fm *models.FileModel) (size int, err error) {
	bname, dname := gsc.separateBucketNameAndFileName(fm.GetPath())
	bucket := gsc.client.Bucket(bname)
	objname := fm.GetName()
	if dname != "" {
		objname = dname + "/" + objname
	}
	obj := bucket.Object(objname)
	w := obj.NewWriter(gsc.ctx)
	if size, err = fmt.Fprintf(w, string(fm.GetData())); err != nil {
		logger.Errorf("Failed to write object: %v", err)
	}
	if err = w.Close(); err != nil {
		logger.Errorf("Failed to close object: %v", err)
	}
	logger.Printf("Upload file '%s'", objname)
	return
}

func (gsc *GoogleStorageClient) Delete(path string) (err error) {
	bname, fname := gsc.separateBucketNameAndFileName(path)
	bucket := gsc.client.Bucket(bname)
	obj := bucket.Object(fname)
	err = obj.Delete(gsc.ctx)
	if err != nil {
		logger.Errorf("Failed to delete file '%s' in %s bucket %s", fname, bname, err.Error())
	}
	return
}

func (gsc *GoogleStorageClient) Download(path string) (fm *models.FileModel, err error) {
	bname, fname := gsc.separateBucketNameAndFileName(path)
	obj := gsc.client.Bucket(bname).Object(fname)
	attr, err := obj.Attrs(gsc.ctx)
	if err != nil {
		logger.Errorf("Failed to download file '%s' in %s bucket %s", fname, bname, err.Error())
		return
	}
	rc, err := obj.NewReader(gsc.ctx)
	if err != nil {
		logger.Errorf("Failed to download file '%s' in %s bucket %s", fname, bname, err.Error())
		return
	}
	defer rc.Close()
	data, err := ioutil.ReadAll(rc)
	fm, _ = models.NewFile(attr.Name, int32(attr.Size), data, attr.Name, attr.Updated, attr.Created, attr.Owner, "")
	logger.Printf("Downloaded file '%s'", fname)
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
			logger.Errorf("Failed to GetBucketList %s", err.Error())
			return
		}
		dirs = append(dirs, battrs.Name)
	}
	logger.Printf("GetBucketList: '%s'", path)
	return
}

func (gsc *GoogleStorageClient) GetFileList(path string) (files []string, err error) {
	var attrs *storage.ObjectAttrs
	var iterr error
	bname, fname := gsc.separateBucketNameAndFileName(path)
	q := &storage.Query{
		Prefix: fname,
	}
	bucket := gsc.client.Bucket(bname)
	it := bucket.Objects(gsc.ctx, q)
	for {
		attrs, iterr = it.Next()
		if iterr == iterator.Done {
			break
		}
		if iterr != nil {
			err = iterr
			logger.Errorf("Failed to GetFileList %s", err.Error())
			return
		}
		s := attrs.Name
		files = append(files, s)
	}
	logger.Printf("GetFileList: '%s'", path)
	return
}

func (gsc *GoogleStorageClient) GetFileInfo(path string) (fm *models.FileModel, err error) {
	bname, fname := gsc.separateBucketNameAndFileName(path)
	bucket := gsc.client.Bucket(bname)
	obj := bucket.Object(fname)
	attr, err := obj.Attrs(gsc.ctx)
	if err != nil {
		logger.Errorf("Failed to GetFileInfo bucket: '%s', file: '%s' %s", bname, fname, err.Error())
		return
	}
	fm, err = models.NewFile(attr.Name, int32(attr.Size), []byte{}, path, attr.Updated, attr.Created, attr.Owner, "")
	logger.Printf("Get File Info file: '%s'", fname)
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
