#!/usr/bin/env bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

function runTest {
  echo -e "Running tests...\n"
  find ./test/* -maxdepth 0 | xargs -I{} bash -c 'echo -e "\n\nRunning {}\n"; bats -p {} || true'
  echo -e "End of tests.\n"
}

function existingFsStop {
  echo -e "Stopping existing fs's\n"
  systemctl --user stop project-fs
  echo -e "Stopped existing fs.\n"
}

function existingFsStart {
  echo -e "Starting original fs\n"
  systemctl --user start project-fs
  echo -e "Original fs started\n"
}

function testFsStop {
  echo -e "Stopping fs under test.\n"
  # Kill last background job (should be test fs)
  kill -9 %-
  sudo umount -l ~/.projectfs/mnt
  echo -e "Stopped test fs.\n"
}

function testFsStart {
  echo -e "Starting test fs..."
  sudo umount -l ~/.projectfs/mnt
  # start projectfs and send to background
  source $DIR/../.env.test && \
    $DIR/../target/projectfs ~/.projectfs/mnt &> $DIR/test-fs.log &
  sleep 3
  echo -e "Test fs started."
}
