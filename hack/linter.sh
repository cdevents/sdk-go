#!/bin/bash

set -e -o pipefail

if [[ "$DISABLE_LINTER" == "true" ]]
then
  exit 0
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if ! [[ -x "$(command -v golangci-lint)" ]]; then
	echo "Installing GolangCI-Lint"
	pushd "${DIR}"
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.12.2
	popd
fi

export GO111MODULE=on
golangci-lint run \
  --timeout 30m \
  --verbose \
  --build-tags testonly
