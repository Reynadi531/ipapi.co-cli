APP_NAME=ipapi
BUILDDIR = bin
VERSION := $(shell git describe --tags --always --dirty)
BUILD := $(shell date +%Y-%m-%d\ %H:%M)
LDFLAGS=-ldflags="-w -s -X 'libcommon.Version=${VERSION}' -X 'libcommon.Build=${BUILD}'"

.PHONY: build

build:
	go build ${LDFLAGS} -o ${BUILDDIR}/${APP_NAME} cmd/main.go