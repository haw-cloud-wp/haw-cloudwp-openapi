package Controller

import (
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/scrapes/haw-cloudwp-openapi/src/commons"
	openapi "github.com/scrapes/haw-cloudwp-openapi/src/go"
	"github.com/scrapes/haw-cloudwp-openapi/src/middleware"
	"github.com/scrapes/haw-cloudwp-openapi/src/storage"
	"github.com/scrapes/haw-cloudwp-openapi/src/utils"
	"log"
	"net/http"
	"os"
)

const gcloudDefaultBucket = "customer_bucker"

type ApiController struct {
}

func (a *ApiController) GetFilesName(ctx context.Context, s string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ApiController) OptionsFilesName(ctx context.Context, s string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ApiController) GetFiles(ctx context.Context) (openapi.ImplResponse, error) {
	gstorage := new(storage.GCloudStorage).Init(new(commons.AllowAllPermission))
	bucket := new(commons.Bucket).Init(gstorage, gcloudDefaultBucket)
	err, files := bucket.GetObjects()

	_, cc := middleware.GetToken(ctx)
	log.Println("HasScope: ", cc.HasScope("access:bucket_ZZZ"))

	if err != nil {
		return openapi.ImplResponse{}, err
	}
	response := openapi.GetFiles200Response{
		Bucket: bucket.GetName(),
		Files: utils.Map(files, func(o commons.IObjectInfo) openapi.FileInfo {
			return openapi.FileInfo{
				Name:    o.GetName(),
				Size:    o.GetSize(),
				Lastmod: o.GetLastMod(),
			}
		}),
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: response,
	}, nil
}

func (a *ApiController) OptionsFiles(ctx context.Context) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ApiController) OptionsFileUpload(ctx context.Context) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: nil,
	}, nil
}

func (a *ApiController) PutFileUpload(ctx context.Context, s string, file *os.File) (openapi.ImplResponse, error) {
	f, err := os.Open(file.Name())
	if err != nil {
		log.Println(err)
		return openapi.ImplResponse{}, err
	}

	gstorage := new(storage.GCloudStorage).Init(new(commons.AllowAllPermission))
	bucket := new(commons.Bucket).Init(gstorage, gcloudDefaultBucket)
	obj := new(commons.Object).Init(bucket, s)
	err = obj.Set(f)

	if err != nil {
		log.Println(err)
		return openapi.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: nil,
		}, err
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: struct {
			Message string
		}{"Successfully Uploaded"},
	}, nil
}
func (a *ApiController) OptionsApiExternal(ctx context.Context) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: nil,
	}, nil
}

func (a *ApiController) PostUser(ctx context.Context, request openapi.PostUserRequest) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: nil,
	}, nil
}

func (a *ApiController) GetUsersUserId(ctx context.Context, i int32) (openapi.ImplResponse, error) {
	user := openapi.User{
		Id:            i,
		FirstName:     "user x",
		LastName:      "yxy",
		Email:         "test@test.de",
		DateOfBirth:   "01.01.1820",
		EmailVerified: false,
		CreateDate:    "now",
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: user,
	}, nil
}

func (a *ApiController) OptionsUsersUserId(ctx context.Context, i int32) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: nil,
	}, nil
}

func (a *ApiController) PatchUsersUserId(ctx context.Context, i int32, request openapi.PatchUsersUserIdRequest) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: nil,
	}, nil
}

func (a *ApiController) OptionsUser(ctx context.Context) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: nil,
	}, nil
}

func (a *ApiController) GetApiExternal(ctx context.Context) (openapi.ImplResponse, error) {
	token := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	rpns := openapi.ImplResponse{}
	_ = token.CustomClaims.(*middleware.CustomClaims)
	log.Printf("api call")
	rpns.Code = http.StatusOK
	rpns.Body = `{"message":"PaPing"}`
	return rpns, nil
}
