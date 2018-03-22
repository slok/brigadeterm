#!/usr/bin/env sh

set -o errexit
set -o nounset


src=./cmd/brigadeterm
out=./bin/brigadeterm
goarch=amd64
platform=${1:-linux}

# Select the release type.
case "${platform}" in
    "linux" )
        echo "Building linux release..."
        goos=linux
        binary_ext=-linux-amd64

    ;;
    "darwin" )
        echo "Building darwin release..."
        goos=linux
        binary_ext=-darwin-amd64
    ;;
    "windows" )
        echo "Building windows release..."
        goos=windows
        binary_ext=-windows-amd64.exe
    ;;
esac

final_out=${out}${binary_ext}
ldf_cmp="-w -extldflags '-static'"
f_ver="-X main.Version=${VERSION:-dev}"

echo "Building binary at ${final_out}"
GOOS=${goos} GOARCH=${goarch} CGO_ENABLED=0 go build -o ${final_out} --ldflags "${ldf_cmp} ${f_ver}"  ${src}