#!/usr/bin/env sh

function build {
    echo "Building binary for" $1
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build $1
}

echo "Testing.."
go get -u github.com/onsi/ginkgo/ginkgo
ginkgo -r 

build cmd/protocol.go
build resource/check/cmd/check.go
build resource/in/cmd/in.go
build resource/out/cmd/out.go
