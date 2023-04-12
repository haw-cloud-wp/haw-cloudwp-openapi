#!/bin/sh
openapi-generator generate \
-i ../reference/customerfacing.yaml \
-g go-server \
-o . \
--additional-properties=outputAsLibrary=true