#!/bin/bash -eux

pushd dp-code-list-api
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4
  make lint
popd