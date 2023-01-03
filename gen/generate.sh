#!/bin/bash

PATH="$PATH:$(go env GOPATH)/bin"
export PATH

rm -rf go/*
mkdir -p go

find ../src/proto -name '*.proto' -print0 | while IFS= read -r -d '' file
do
  echo "Generating protos for $file"
  protoc -I ../src/proto --go_out=go --go-grpc_out=go "$file"
done

mv go/github.com/EmortalMC/grpc-api-specs/gen/go/* go/
