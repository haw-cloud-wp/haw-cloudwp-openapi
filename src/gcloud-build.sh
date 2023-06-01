#!/bin/sh

docker build . --platform linux/arm64 --tag cloudwp.azurecr.io/openapi:arm
docker build . --platform linux/amd64 --tag gcr.io/subtle-digit-382316/openapi
docker push gcr.io/subtle-digit-382316/openapi
docker tag gcr.io/subtle-digit-382316/openapi cloudwp.azurecr.io/openapi:latest
docker push cloudwp.azurecr.io/openapi:latest
docker push cloudwp.azurecr.io/openapi:arm