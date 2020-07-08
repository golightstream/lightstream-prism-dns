#!/bin/sh

build() {
  _BINARY=lightstream-prism-dns-$1-$2
  env GOOS=$1 GOARCH=$2 go build -o $_BINARY
  upx $_BINARY
  if [ "$1" = "windows" ]; then
    mv $_BINARY $_BINARY.exe
  fi
}


build linux 386
build linux amd64
build windows 386
build windows amd64

