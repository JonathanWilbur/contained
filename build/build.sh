#!/bin/sh -e
# export GO_EXTLINK_ENABLED=0
export CGO_ENABLED=0
# export GOOS=darwin
# export GOARCH=amd64
go build -ldflags=-s -o ./dist/contained ./source/main.go