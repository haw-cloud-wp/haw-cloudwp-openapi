package service

import (
	"context"
	"encoding/binary"
	"errors"
	"github.com/scrapes/haw-cloudwp-openapi/src/commons"
	"github.com/scrapes/haw-cloudwp-openapi/src/db"
	"github.com/scrapes/haw-cloudwp-openapi/src/middleware"
	"github.com/scrapes/haw-cloudwp-openapi/src/storage"
	openapi "github.com/scrapes/haw-cloudwp-openapi/src/v1/go"
	"net/http"
	"os"
	"strconv"
)

func GetInternalServerError(err error) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Code: http.StatusInternalServerError,
		Body: struct {
			Error string
		}{Error: err.Error()},
	}, err
}

type V1Service struct {
	db *db.Connection
}

func (v *V1Service) getOrCreateUser(userID string) *db.User {
	dbUser := new(db.User)
	rows := v.db.DB.Model(dbUser).Preload("Access").First(dbUser, "auth0_id = ?", userID)
	if !(rows.RowsAffected > 0) {
		dbUser.Auth0ID = userID
		dbUser.Access = make([]db.Bucket, 0)
		v.db.DB.Create(dbUser)
	}
	return dbUser
}

func (v *V1Service) SetDB(db *db.Connection) {
	v.db = db
}

func (v *V1Service) DeleteV1BucketName(ctx context.Context, bucketName string) (openapi.ImplResponse, error) {
	_, cc := middleware.GetToken(ctx)
	permission := new(commons.ClaimsPermissionHandler).Init(cc)
	gStorage := new(storage.GCloudStorage).Init(permission)
	bucketToDelete := new(commons.Bucket).Init(gStorage, bucketName)
	err := bucketToDelete.Delete()

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

func (v *V1Service) DeleteV1FileName(ctx context.Context, bucketName string, fileName string) (openapi.ImplResponse, error) {
	_, cc := middleware.GetToken(ctx)
	perm := new(commons.ClaimsPermissionHandler).Init(cc)
	gStorage := new(storage.GCloudStorage).Init(perm)
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
	permission := new(commons.ClaimsPermissionHandler).Init(cc)
	gStorage := new(storage.GCloudStorage).Init(permission)
	bucket := new(commons.Bucket).Init(gStorage, bucketName)
	name := bucket.GetName()

	if binary.Size(name) <= 0 { //ist das einproblem wenn der Bucket "" heißt?
		return GetInternalServerError(errors.New("bucket name empty"))
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.BucketInfo{
			Owner:     "kommt noch aus der SQLDB111979",
			CreatedAt: "kommt noch aus der SQLDB111979",
		}, //wo wird der Name zurückgeben #anton hiiiilfe
	}, nil
}

func (v *V1Service) GetV1Buckets(ctx context.Context) (openapi.ImplResponse, error) {
	_, cc := middleware.GetToken(ctx)
	user := v.getOrCreateUser(cc.UserID)
	var response []openapi.Bucket
	for _, bucket := range user.Access {
		response = append(response, openapi.Bucket{
			Id:   strconv.Itoa(int(bucket.ID)),
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
	if binary.Size(bucketName) <= 0 { //ist das einproblem wenn der Bucket "" heißt?
		return GetInternalServerError(errors.New("bucket name empty"))
	}

	_, cc := middleware.GetToken(ctx)
	permission := new(commons.ClaimsPermissionHandler).Init(cc)
	gStorage := new(storage.GCloudStorage).Init(permission)
	bucket := new(commons.Bucket).Init(gStorage, bucketName)
	_, files := bucket.GetObjects()

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
	gstore := new(storage.GCloudStorage).Init(new(commons.ClaimsPermissionHandler).Init(cc))
	err, _ := gstore.CreateBucket(request.Name)
	if err != nil {
		return GetInternalServerError(err)
	}

	dbBucket := new(db.Bucket)
	dbBucket.Name = request.Name
	v.db.DB.Create(dbBucket)

	dbUser := v.getOrCreateUser(cc.UserID)
	dbUser.Access = append(dbUser.Access, *dbBucket)
	v.db.DB.Updates(dbUser)

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: struct {
			Message string
		}{Message: "OK"},
	}, nil
}

func (v *V1Service) PutV1FileName(ctx context.Context, s string, s2 string, file *os.File) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}
