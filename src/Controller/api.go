package Controller

import (
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	openapi "github.com/scrapes/haw-cloudwp-openapi/src/go"
	"github.com/scrapes/haw-cloudwp-openapi/src/middleware"
	"github.com/scrapes/haw-cloudwp-openapi/src/storage"
	"log"
	"net/http"
	"os"
)

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
	files, err := storage.GcloudListFiles()
	if err != nil {
		return openapi.ImplResponse{}, err
	}
	response := openapi.GetFiles200Response{
		Bucket: storage.GcloudBucketName,
		Files:  files,
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
		return openapi.ImplResponse{}, err
	}

	err = storage.GcloudUploadFile(f, s)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: nil,
		}, err
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: nil,
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