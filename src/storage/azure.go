package storage

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/scrapes/haw-cloudwp-openapi/src/commons"
	"io"
	"log"
	"os"
	"time"
)

const (
	URL = "https://cloudwp.blob.core.windows.net" //@TODO
)

type AzureStorage struct {
	credential  *azidentity.DefaultAzureCredential
	permissions commons.IPermission
	ctx         context.Context
	client      *azblob.Client
}

func (A *AzureStorage) Init(permissions commons.IPermission) commons.IStorage {
	A.credential, _ = azidentity.NewDefaultAzureCredential(nil)
	A.permissions = permissions
	A.ctx = context.Background()
	var err error
	A.client, err = azblob.NewClient(URL, A.credential, nil)
	if err != nil {
		fmt.Print(fmt.Errorf("storage.NewClient: %v", err))
	}
	return A
}

func (A *AzureStorage) Close() error {
	return nil
}

func (A *AzureStorage) GetObjectSize(bucket commons.IBucket, object commons.IObject) (error, int64) {
	//TODO implement me
	panic("implement me")
}

func (A *AzureStorage) GetObjectLastMod(bucket commons.IBucket, object commons.IObject) (error, time.Time) {
	//TODO implement me
	panic("implement me")
}

func (A *AzureStorage) DeleteBucket(bucket commons.IBucket) error {
	//return G.client.Bucket(GCLOUD_PROJECT_ID + "__" + bucket.GetName()).Delete(G.ctx)
	_, err := A.client.DeleteContainer(A.ctx, bucket.GetName(), nil)
	return err
}

func (A *AzureStorage) CreateBucket(name string) (error, commons.IBucket) {
	_, err := A.client.CreateContainer(A.ctx, name, nil)
	return err, nil
}

func (A *AzureStorage) GetObjects(bucket commons.IBucket) (error, []commons.IObjectInfo) {
	ctx, cancel := context.WithTimeout(A.ctx, time.Second*10)
	defer cancel()
	var objects []commons.IObjectInfo
	pager := A.client.NewListBlobsFlatPager(bucket.GetName(), nil)
	for {
		obj, err := pager.NextPage(ctx)
		if err != nil {
			log.Println("Error on pager: ", err.Error())
			break
		}

		for _, item := range obj.ListBlobsFlatSegmentResponse.Segment.BlobItems {
			objects = append(objects, new(commons.ObjectInfo).Init(*item.Name, *item.Properties.ContentLength, *item.Properties.LastModified))
		}
		if !pager.More() {
			break
		}
	}
	return nil, objects
}

func (A *AzureStorage) DeleteObject(bucket commons.IBucket, object commons.IObject) error {
	_, err := A.client.DeleteBlob(A.ctx, bucket.GetName(), object.GetName(), nil)
	return err
}

func (A *AzureStorage) GetObjectStream(bucket commons.IBucket, object commons.IObject) (error, *io.Reader) {
	//TODO implement me
	panic("implement me")
}

func (A *AzureStorage) GetObject(bucket commons.IBucket, object commons.IObject) (error, *os.File) {
	temp, err := os.CreateTemp("", "obj-")

	if err != nil {
		return err, nil
	}

	_, err = A.client.DownloadFile(A.ctx, bucket.GetName(), object.GetName(), temp, nil)
	if err != nil {
		return err, nil
	}

	return nil, temp
}

func (A *AzureStorage) SetObject(bucket commons.IBucket, object commons.IObject, data *os.File) error {
	_, err := A.client.UploadFile(A.ctx, bucket.GetName(), object.GetName(), data, nil)
	return err
}

func (A *AzureStorage) GetPermission(permission string) bool {
	return A.permissions.HasPermission(permission)
}
