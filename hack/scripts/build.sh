#!/usr/bin/env sh

set -o errexit
set -o nounset


SRC=./cmd/brigadeterm
OUT=./bin/brigadeterm
LDF_CMP="-w -extldflags '-static'"

echo "Building binary at ${OUT}"
CGO_ENABLED=0 go build -o ${OUT} --ldflags "${LDF_CMP}"  ${SRC}