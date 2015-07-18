#!/usr/bin/env bash

source fs-control.sh

existingFsStop
testFsStart
runTest
testFsStop
existingFsStart
