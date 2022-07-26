.PHONY: clean build image

BIN_NAME := payslip
MAIN_DIRECTORY := ./cmd/payslip

TAG_NAME := $(shell git tag -l --contains HEAD)

# Default build target
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
DOCKER_BUILD_PLATFORMS ?= linux/amd64,linux/arm64

default: clean build

clean:
	rm -rf cover.out

dist:
	mkdir dist

build: clean dist
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -v -o "./dist/${GOOS}/${GOARCH}/${BIN_NAME}" ${MAIN_DIRECTORY}

build-linux-arm64: export GOOS := linux
build-linux-arm64: export GOARCH := arm64
build-linux-arm64:
	make build

build-linux-amd64: export GOOS := linux
build-linux-amd64: export GOARCH := amd64
build-linux-amd64:
	make build

## Build Multi archs Docker image
multi-arch-image-%: build-linux-amd64 build-linux-arm64
	docker buildx build $(DOCKER_BUILDX_ARGS) --progress=chain -t moutoum/$(BIN_NAME):$* --platform=$(DOCKER_BUILD_PLATFORMS) -f Dockerfile .
