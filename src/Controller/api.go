package Controller

import (
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	openapi "github.com/scrapes/haw-cloudwp-openapi/src/go"
	"github.com/scrapes/haw-cloudwp-openapi/src/middleware"
	"log"
	"net/http"
	"reflect"
)

type ApiController struct {
}

func (a *ApiController) OptionsUser(ctx context.Context) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ApiController) OptionsUsersUserId(ctx context.Context, i interface{}) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
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

func (a *ApiController) GetUsersUserId(ctx context.Context, i interface{}) (openapi.ImplResponse, error) {
	log.Printf(reflect.TypeOf(i).String())
	return openapi.ImplResponse{Code: http.StatusOK, Body: `{"message":"PaPing"}`}, nil
}

func (a *ApiController) OptionsApiExternal(ctx context.Context) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ApiController) PatchUsersUserId(ctx context.Context, i interface{}, request openapi.PatchUsersUserIdRequest) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ApiController) PostUser(ctx context.Context, request openapi.PostUserRequest) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}
