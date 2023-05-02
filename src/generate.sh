#!/bin/sh
openapi-generator generate \
-i ../reference/v1api.yaml \
-g go-server \
-o ./v1 \
--additional-properties=outputAsLibrary=true