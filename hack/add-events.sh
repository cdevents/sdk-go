#!/usr/bin/env bash

BASE_DIR="$( cd "$( dirname "$0" )" >/dev/null 2>&1 && pwd )"
cd $BASE_DIR

./add-event.sh repository created
./add-event.sh repository modified
./add-event.sh repository deleted
./add-event.sh branch created
./add-event.sh branch deleted
./add-event.sh change created
./add-event.sh change updated
./add-event.sh change reviewed
./add-event.sh change merged
./add-event.sh change abandoned
./add-event.sh build started
./add-event.sh build queued
./add-event.sh build finished
./add-event.sh testCase started
./add-event.sh testCase queued
./add-event.sh testCase finished
./add-event.sh testSuite started
./add-event.sh testSuite queued
./add-event.sh testSuite finished
./add-event.sh artifact packaged
./add-event.sh artifact published
./add-event.sh environment created
./add-event.sh environment modified
./add-event.sh environment deleted
./add-event.sh service deployed
./add-event.sh service upgraded
./add-event.sh service rolledback
./add-event.sh service removed
./add-event.sh service published