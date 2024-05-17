#!/bin/bash

set -e

REPO_PATH="$(cd "$(dirname "$0")/.."; pwd)"

RED='\033[0;31m'
GREEN='\033[0;32m'
RESET='\033[0m'

if ! [[ -x "$(command -v goimports)" ]]; then
	echo "Installing goimports"
    go get golang.org/x/tools/cmd/goimports
fi

echo "Running validation scripts..."

function cleanup() {
    # Cleanup
    find . -name 'zz_*' -exec rm {} +
}

function generate() {
    # Generate code
    go run tools/generator.go --resources "${REPO_PATH}"
}

function check() {
    # Check if code was up to date
    GIT_STATUS=$(git status -s)
    if [[ ! -z "${GIT_STATUS}" ]]; then
        echo -e "Changes detected:\n$GIT_STATUS"
        return 1
    fi
}

scripts=(
    cleanup
    generate
    check
)

fail=0
for s in "${scripts[@]}"; do
    echo "RUN ${s}"
    set +e
    ${s}
    result=$?
    set -e
    if [[ "${result}"  -eq 0 ]]; then
        echo -e "${GREEN}PASSED${RESET} ${s}"
    else
        echo -e "${RED}FAILED${RESET} ${s}"
        fail=1
    fi
done
exit "${fail}"