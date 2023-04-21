package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/scrapes/haw-cloudwp-openapi/src/commons"
	"google.golang.org/api/iterator"
	"io"
	"log"
	"os"
	"time"
)

type GCloudStorage struct {
	permissions commons.IPermission
}

func (G *GCloudStorage) Init(permissions commons.IPermission) commons.IStorage {
	G.permissions = permissions
	return G
}

func (G *GCloudStorage) DeleteBucket(bucket commons.IBucket) error {
	//TODO implement me
	panic("implement me")
}

func (G *GCloudStorage) CreateBucket(name string) (error, commons.IBucket) {
	//TODO implement me
	panic("implement me")
}

func (G *GCloudStorage) GetObjects(bucket commons.IBucket) (error, []commons.IObject) {
	// bucket := "bucket-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err), nil
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	it := client.Bucket(bucket.GetName()).Objects(ctx, nil)
	var objects []commons.IObject
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects: %v", bucket.GetName(), err), nil
		}
		objects = append(objects, new(commons.Object).Init(bucket, attrs.Name))
	}
	return nil, objects
}

func (G *GCloudStorage) DeleteObject(bucket commons.IBucket, object commons.IObject) error {
	//TODO implement me
	panic("implement me")
}

func (G *GCloudStorage) GetObjectStream(bucket commons.IBucket, object commons.IObject) (error, *io.Reader) {
	//TODO implement me
	panic("implement me")
}

func (G *GCloudStorage) GetObject(bucket commons.IBucket, object commons.IObject) (error, *os.File) {
	//TODO implement me
	panic("implement me")
}

func (G *GCloudStorage) SetObject(bucket commons.IBucket, object commons.IObject, data *os.File) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := client.Bucket(bucket.GetName()).Object(object.GetName())

	o = o.If(storage.Conditions{DoesNotExist: false})

	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, data); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	log.Printf("Blob %v uploaded.\n", object)
	return nil
}

func (G *GCloudStorage) GetPermission(permission string) bool {
	return G.permissions.HasPermission(permission)
}
