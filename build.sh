#!/bin/bash

export GOOS=windows
export CGO_ENABLED=1
export CC="x86_64-w64-mingw32-gcc"
go build -v -o main-test.exe main.go
