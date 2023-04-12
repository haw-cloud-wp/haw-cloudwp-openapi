package Controller

import (
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	openapi "github.com/scrapes/haw-cloudwp-openapi/src/go"
	"github.com/scrapes/haw-cloudwp-openapi/src/middleware"
	"net/http"
)

type ApiController struct {
	openapi.DefaultApiServicer
}

func (a *ApiController) GetApiExternal(ctx context.Context) (openapi.ImplResponse, error) {
	token := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	rpns := openapi.ImplResponse{}
	claims := token.CustomClaims.(*middleware.CustomClaims)

	if !claims.HasScope("read:messages") {
		rpns.Code = http.StatusForbidden
		rpns.Body = `{"message":"Insufficient scope."}`
		return rpns, nil
	}

	rpns.Code = http.StatusOK
	rpns.Body = `{"message":"PaPing"}`
	return rpns, nil
}

func (a *ApiController) GetUsersUserId(ctx context.Context, i interface{}) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
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
