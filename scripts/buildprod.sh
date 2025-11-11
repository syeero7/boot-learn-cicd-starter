#!/bin/bash

set +e

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notely 2>error.log
EXIT_CODE=$?

if [ $EXIT_CODE -ne 0 ]; then
  echo "BUILD_FAILED_MARKER" >>error.log
  exit $EXIT_CODE
fi
