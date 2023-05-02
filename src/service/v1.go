package service

import (
	"context"
	openapi "github.com/scrapes/haw-cloudwp-openapi/src/v1/go"
	"net/http"
	"os"
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
}

func (v *V1Service) DeleteV1BucketName(ctx context.Context, s string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) DeleteV1FileName(ctx context.Context, s string, s2 string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) GetV1BucketName(ctx context.Context, s string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) GetV1Buckets(ctx context.Context) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) GetV1FileName(ctx context.Context, s string, s2 string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) GetV1Files(ctx context.Context, s string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
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

func (v *V1Service) PostV1BucketName(ctx context.Context, s string, request openapi.PostV1BucketNameRequest) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (v *V1Service) PutV1FileName(ctx context.Context, s string, s2 string, file *os.File) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}
