#!/usr/bin/env sh

set -o errexit
set -o nounset

SRC=./cmd/brigadeterm
OUT=./bin/brigadeterm
LDF_CMP="-w -extldflags '-static'"
F_VER="-X main.Version=${VERSION:-dev}"

echo "Building binary at ${OUT}"
CGO_ENABLED=0 go build -o ${OUT} --ldflags "${LDF_CMP} ${F_VER}"  ${SRC}