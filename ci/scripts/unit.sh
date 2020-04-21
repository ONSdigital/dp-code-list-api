#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-code-list-api
  make test
popd
