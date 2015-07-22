#!/usr/bin/env bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
source $DIR/fs-control.sh

existingFsStop
testFsStart
runTest
testFsStop
existingFsStart
