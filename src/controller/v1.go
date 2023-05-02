package controller

import (
	"github.com/gorilla/mux"
	"github.com/scrapes/haw-cloudwp-openapi/src/commons"
	"github.com/scrapes/haw-cloudwp-openapi/src/middleware"
	"github.com/scrapes/haw-cloudwp-openapi/src/storage"
	openapi "github.com/scrapes/haw-cloudwp-openapi/src/v1/go"
	"net/http"
	"os"
)

type V1Controller struct {
	openapi.DefaultApiController
	ErrHandler openapi.ErrorHandler
	Service    openapi.DefaultApiServicer
}

func (c *V1Controller) Init(service openapi.DefaultApiServicer) *V1Controller {
	c.ErrHandler = openapi.DefaultErrorHandler
	c.Service = service
	return c
}

func (c *V1Controller) GetV1FileName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bucketNameParam := params["BucketName"]

	fileNameParam := params["FileName"]
	_, cc := middleware.GetToken(r.Context())
	cstore := new(storage.GCloudStorage).Init(new(commons.AllowAllPermission).Init(cc))
	bucket := new(commons.Bucket).Init(cstore, bucketNameParam)
	obj := new(commons.Object).Init(bucket, fileNameParam)

	err, dataFile := obj.Get()

	if err != nil {
		c.ErrHandler(w, r, err, &openapi.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: err.Error(),
		})
		return
	}

	http.ServeFile(w, r, dataFile.Name())
	os.Remove(dataFile.Name())
}
