#!/bin/bash

set -e -o pipefail

if [[ "$DISABLE_LINTER" == "true" ]]
then
  exit 0
fi

export GO111MODULE=on
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15 run \
  --timeout 30m \
  --verbose \
  --build-tags testonly
