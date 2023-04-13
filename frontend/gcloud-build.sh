#!/bin/sh

docker build . --platform linux/amd64 --tag gcr.io/subtle-digit-382316/frontend
docker push gcr.io/subtle-digit-382316/frontend