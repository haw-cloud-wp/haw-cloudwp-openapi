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
	ctx         context.Context
	client      *storage.Client
}

func (G *GCloudStorage) Init(permissions commons.IPermission) commons.IStorage {
	G.permissions = permissions
	G.ctx = context.Background()
	var err error
	G.client, err = storage.NewClient(G.ctx)
	if err != nil {
		fmt.Print(fmt.Errorf("storage.NewClient: %v", err))
	}
	return G
}

func (G *GCloudStorage) Close() error {
	err := G.client.Close()
	if err != nil {
		return err
	}

	return nil
}

func (G *GCloudStorage) GetObjectSize(bucket commons.IBucket, object commons.IObject) (error, int64) {
	//TODO implement me
	panic("implement me")
}

func (G *GCloudStorage) GetObjectLastMod(bucket commons.IBucket, object commons.IObject) (error, time.Time) {
	//TODO implement me
	panic("implement me")
}

func (G *GCloudStorage) DeleteBucket(bucket commons.IBucket) error {
	//TODO implement me
	panic("implement me")
}

func (G *GCloudStorage) CreateBucket(name string) (error, commons.IBucket) {
	//TODO implement me
	panic("implement me")
}

func (G *GCloudStorage) GetObjects(bucket commons.IBucket) (error, []commons.IObjectInfo) {
	// bucket := "bucket-name"
	ctx, cancel := context.WithTimeout(G.ctx, time.Second*10)
	defer cancel()

	it := G.client.Bucket(bucket.GetName()).Objects(ctx, nil)
	var objects []commons.IObjectInfo
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects: %v", bucket.GetName(), err), nil
		}
		objects = append(objects, new(commons.ObjectInfo).Init(attrs.Name, attrs.Size, attrs.Updated))
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

	o = o.If(storage.Conditions{DoesNotExist: true})

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
