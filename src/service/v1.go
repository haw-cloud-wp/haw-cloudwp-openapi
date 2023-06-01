package service

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/scrapes/haw-cloudwp-openapi/src/commons"
	"github.com/scrapes/haw-cloudwp-openapi/src/db"
	"github.com/scrapes/haw-cloudwp-openapi/src/middleware"
	"github.com/scrapes/haw-cloudwp-openapi/src/storage"
	openapi "github.com/scrapes/haw-cloudwp-openapi/src/v1/go"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "cloud.google.com/go/translate/apiv3/translatepb"
)

var APIMKEY = ""

func GetInternalServerError(err error) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Code: http.StatusInternalServerError,
		Body: struct {
			Error string
		}{Error: err.Error()},
	}, err
}

func GetNotFound() (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Code: http.StatusNotFound,
		Body: struct {
			Error string
		}{Error: "Object not found"},
	}, nil
}

type V1Service struct {
	db      *db.Connection
	storage commons.IStorage
	isAzure bool
}

func (v *V1Service) Init(db *db.Connection, store commons.IStorage, isAzure bool) *V1Service {
	v.db = db
	v.storage = store
	v.isAzure = isAzure
	return v
}

func (v *V1Service) GetV1BucketBucketNameTranslateFileName(ctxs context.Context, s string, s2 string) (openapi.ImplResponse, error) {
	if !v.isAzure {
		ctx := context.Background()
		// This snippet has been automatically generated and should be regarded as a code template only.
		// It will require modifications to work:
		// - It may require correct/in-range values for request initialization.
		// - It may require specifying regional endpoints when creating the service client as shown in:
		//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
		c, err := translate.NewTranslationClient(ctx)
		if err != nil {
			// TODO: Handle error.
		}
		defer c.Close()

		req := &translatepb.TranslateDocumentRequest{
			Parent:             fmt.Sprintf("projects/%s/locations/global", "664861925166"),
			SourceLanguageCode: "de",
			TargetLanguageCode: "en",
			DocumentInputConfig: &translatepb.DocumentInputConfig{
				Source: &translatepb.DocumentInputConfig_GcsSource{
					GcsSource: &translatepb.GcsSource{
						InputUri: fmt.Sprintf("gs://%s/%s", storage.GCLOUD_PROJECT_ID+"__"+s, s2),
					},
				},
			},
			DocumentOutputConfig: &translatepb.DocumentOutputConfig{
				Destination: &translatepb.DocumentOutputConfig_GcsDestination{
					GcsDestination: &translatepb.GcsDestination{
						OutputUriPrefix: fmt.Sprintf("gs://%s/", storage.GCLOUD_PROJECT_ID+"__"+s),
					},
				},
			},
		}
		resp, err := c.TranslateDocument(ctx, req)
		if err != nil {
			return GetInternalServerError(err)
		}
		info := resp.String()
		return openapi.ImplResponse{
			Code: http.StatusOK,
			Body: struct {
				Message string
				Info    string
			}{Message: "OK", Info: info},
		}, nil
	} else {
		accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"), os.Getenv("AZURE_STORAGE_ACCOUNT_KEY")
		var containerName = s

		credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
		if err != nil {
			log.Println("Translate err: ", err.Error())
			return GetInternalServerError(err)
		}

		sourceSASParams, err := sas.BlobSignatureValues{
			Protocol:      sas.ProtocolHTTPS,
			StartTime:     time.Now().UTC(),
			ExpiryTime:    time.Now().UTC().Add(1 * time.Hour),
			Permissions:   to.Ptr(sas.BlobPermissions{Read: true, Create: true, Write: true, Tag: true}).String(),
			ContainerName: containerName,
			BlobName:      s2,
		}.SignWithSharedKey(credential)
		if err != nil {
			log.Println("Translate err: ", err.Error())
			return GetInternalServerError(err)
		}
		sourcesasURL := fmt.Sprintf("https://%s.blob.core.windows.net/?%s", accountName, sourceSASParams.Encode())

		targetSASParams, err := sas.BlobSignatureValues{
			Protocol:      sas.ProtocolHTTPS,
			StartTime:     time.Now().UTC(),
			ExpiryTime:    time.Now().UTC().Add(1 * time.Hour),
			Permissions:   to.Ptr(sas.BlobPermissions{Read: true, Create: true, Write: true, Tag: true}).String(),
			ContainerName: containerName,
			BlobName:      "FR_" + s2,
		}.SignWithSharedKey(credential)
		if err != nil {
			log.Println("Translate err: ", err.Error())
			return GetInternalServerError(err)
		}
		targetsasURL := fmt.Sprintf("https://%s.blob.core.windows.net/?%s", accountName, targetSASParams.Encode())

		httpposturl := "https://cloudwp2.cognitiveservices.azure.com/translator/text/batch/v1.0/batches"
		fmt.Println("Response", httpposturl)

		var jsonData = []byte(fmt.Sprintf(`{
        "inputs": [
            {
                "source": {
                    "sourceUrl": "%s"
                },
                "targets": [
                    {
                        "%s",
                        "language": "fr"
                    }
                ]
            }
        ]
    }`, sourcesasURL, targetsasURL))

		request, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Ocp-Apim-Subscription-Key", APIMKEY)

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return GetInternalServerError(err)
		}
		defer response.Body.Close()
		content, err := io.ReadAll(response.Body)
		if err != nil {
			return GetInternalServerError(err)
		}
		if response.StatusCode >= 300 {
			return GetInternalServerError(errors.New(string(content)))
		}
		return openapi.ImplResponse{
			Code: http.StatusOK,
			Body: struct {
			}{},
		}, nil
	}
}

func (v *V1Service) OptionsV1BucketBucketNameTranslateFileName(ctx context.Context, s string, s2 string, request openapi.OptionsV1BucketBucketNameTranslateFileNameRequest) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) getOrCreateUser(userID string) *db.User {
	dbUser := new(db.User)
	rows := v.db.DB.Model(dbUser).Preload("Access").First(dbUser, "auth0_id = ?", userID)
	if !(rows.RowsAffected > 0) {
		dbUser.Auth0ID = userID
		dbUser.Access = make([]*db.Bucket, 0)
		v.db.DB.Create(dbUser)
	}
	return dbUser
}

func (v *V1Service) SetDB(db *db.Connection) {
	v.db = db
}

func (v *V1Service) SetStorage(store commons.IStorage) {
	v.storage = store
}

func (v *V1Service) DeleteV1BucketName(ctx context.Context, bucketName string) (openapi.ImplResponse, error) {
	_, cc := middleware.GetToken(ctx)
	permission := new(commons.AllowAllPermission).Init(cc)
	gStorage := v.storage.Init(permission)
	bucketToDelete := new(commons.Bucket).Init(gStorage, bucketName)
	err := bucketToDelete.Delete()

	if err != nil {
		return GetInternalServerError(err)
	}

	dbBucket := new(db.Bucket)
	v.db.DB.First(dbBucket, "name = ?", bucketName)
	err = v.db.DB.Model(dbBucket).Association("User").Clear()
	v.db.DB.Updates(dbBucket)
	v.db.DB.Delete(dbBucket)

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: struct {
			Message string
		}{Message: "OK"},
	}, nil
}

func (v *V1Service) DeleteV1FileName(ctx context.Context, bucketName string, fileName string) (openapi.ImplResponse, error) {
	_, cc := middleware.GetToken(ctx)
	perm := new(commons.AllowAllPermission).Init(cc)
	gStorage := v.storage.Init(perm)
	bucket := new(commons.Bucket).Init(gStorage, bucketName)
	objectToDelete := new(commons.Object).Init(bucket, fileName)
	err := objectToDelete.Delete()

	if err != nil {
		return GetInternalServerError(err)
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: struct {
			Message string
		}{Message: "OK"},
	}, nil

}

func (v *V1Service) GetV1BucketName(ctx context.Context, bucketName string) (openapi.ImplResponse, error) {
	_, cc := middleware.GetToken(ctx)
	permission := new(commons.AllowAllPermission).Init(cc)
	gStorage := v.storage.Init(permission)
	bucket := new(commons.Bucket).Init(gStorage, bucketName)
	name := bucket.GetName()

	if binary.Size(name) <= 0 { //ist das einproblem wenn der Bucket "" heißt?
		return GetInternalServerError(errors.New("bucket name empty"))
	}

	dbBucket := new(db.Bucket)
	row := v.db.DB.First(dbBucket, "name = ?", bucketName)
	if row.RowsAffected > 0 {
		return openapi.ImplResponse{
			Code: http.StatusOK,
			Body: openapi.BucketInfo{
				Owner:     "",
				CreatedAt: dbBucket.CreatedAt,
			},
		}, nil
	}
	return GetNotFound()
}

func (v *V1Service) GetV1Buckets(ctx context.Context) (openapi.ImplResponse, error) {
	_, cc := middleware.GetToken(ctx)
	user := v.getOrCreateUser(cc.UserID)
	var response []openapi.Bucket
	for _, bucket := range user.Access {
		response = append(response, openapi.Bucket{
			Id:   bucket.Name,
			Name: bucket.Name,
		})
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: response,
	}, nil
}

func (v *V1Service) GetV1FileName(ctx context.Context, s string, s2 string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) GetV1Files(ctx context.Context, bucketName string) (openapi.ImplResponse, error) {
	_, cc := middleware.GetToken(ctx)
	log.Println(bucketName)
	permission := new(commons.AllowAllPermission).Init(cc)
	gStorage := v.storage.Init(permission)
	bucket := new(commons.Bucket).Init(gStorage, bucketName)
	err, files := bucket.GetObjects()
	if err != nil {
		return GetInternalServerError(err)
	}

	var fileInfos []openapi.FileInfo
	for _, file := range files {
		fileInfos = append(fileInfos, openapi.FileInfo{
			File: openapi.File{
				Name: file.GetName(),
			},
			Size:    file.GetSize(),
			Lastmod: file.GetLastMod(),
		})
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: fileInfos,
	}, nil

}

func (v *V1Service) OptionsV1BucketName(ctx context.Context, s string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) OptionsV1Buckets(ctx context.Context) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) OptionsV1FileName(ctx context.Context, s string, s2 string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) OptionsV1Files(ctx context.Context, s string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) PatchV1BucketName(ctx context.Context, s string, permissions []openapi.Permission) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

// Create Bucket
func (v *V1Service) PostV1BucketName(ctx context.Context, s string, request openapi.PostV1BucketNameRequest) (openapi.ImplResponse, error) {
	_, cc := middleware.GetToken(ctx)
	gstore := v.storage.Init(new(commons.AllowAllPermission).Init(cc))
	err, _ := gstore.CreateBucket(request.Name)
	if err != nil {
		return GetInternalServerError(err)
	}

	dbBucket := new(db.Bucket)
	dbBucket.Name = request.Name
	v.db.DB.Create(dbBucket)

	dbUser := v.getOrCreateUser(cc.UserID)
	dbUser.Access = append(dbUser.Access, dbBucket)
	v.db.DB.Updates(dbUser)

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: struct {
			Message string
		}{Message: "OK"},
	}, nil
}

func (v *V1Service) PutV1FileName(ctx context.Context, s string, s2 string, file *os.File) (openapi.ImplResponse, error) {
	_, cc := middleware.GetToken(ctx)
	store := v.storage.Init(new(commons.AllowAllPermission).Init(cc))
	bucket := new(commons.Bucket).Init(store, s)
	obj := new(commons.Object).Init(bucket, s2)
	f, err := os.Open(file.Name())
	if err != nil {
		return GetInternalServerError(err)
	}
	err = obj.Set(f)
	if err != nil {
		return GetInternalServerError(err)
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: struct {
			Message string
		}{Message: "OK"},
	}, nil
}
