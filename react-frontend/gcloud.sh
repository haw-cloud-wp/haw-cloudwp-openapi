#!/bin/sh
docker build . --platform linux/arm64 --tag gcr.io/subtle-digit-382316/frontend:arm
docker build . --platform linux/amd64 --tag gcr.io/subtle-digit-382316/frontend
docker push gcr.io/subtle-digit-382316/frontend
docker tag gcr.io/subtle-digit-382316/frontend cloudwp.azurecr.io/frontend:latest
docker push cloudwp.azurecr.io/frontend:latest
