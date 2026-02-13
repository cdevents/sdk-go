#!/bin/bash

set -e

REPO_PATH="$(cd "$(dirname "$0")/.."; pwd)"

RED='\033[0;31m'
GREEN='\033[0;32m'
RESET='\033[0m'

function apply_patches() {
    # Apply patches to fix upstream schema bugs
    PATCHES_DIR="${REPO_PATH}/hack/patches"
    if [[ -d "${PATCHES_DIR}" ]]; then
        echo "Applying patches from ${PATCHES_DIR}..."
        for patch in "${PATCHES_DIR}"/*.patch; do
            if [[ -f "${patch}" ]]; then
                echo "  Applying $(basename ${patch})..."
                # Apply patch from repo root, ignore whitespace, and don't fail if already applied
                (cd "${REPO_PATH}" && patch -p1 -N -r - < "${patch}") || true
            fi
        done
    fi
}

function revert_patches() {
    # Revert patches from submodules to keep them clean
    echo "Reverting patches from submodules..."
    git submodule foreach 'git reset --hard HEAD && git clean -fd' > /dev/null 2>&1 || true
}

scripts=(
    apply_patches
    revert_patches
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
