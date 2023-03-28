#!/bin/sh

build() {
  _BINARY=lightstream-prism-dns-$1-$2
  env GOOS=$1 GOARCH=$2 make BINARY="$_BINARY"

  if [ "$1" = "windows" ]; then
    mv $_BINARY $_BINARY.exe
  fi
}


build linux amd64
build linux 386
build windows amd64
build darwin amd64
build darwin arm64

