BUILD_DIR ?= $(dir $(realpath -s $(firstword $(MAKEFILE_LIST))))/build
VERSION ?= $(shell git describe --tags --always)
GOOS ?= $(shell uname | tr '[:upper:]' '[:lower:]')
GOARCH ?= $(shell [[ `uname -m` = "x86_64" ]] && echo "amd64" || echo "arm64" )
GOPROXY ?= "https://proxy.golang.org,direct"
RELEASE_REPO ?= public.ecr.aws/brandonwagner
RELEASE_PLATFORM ?= "--platform=linux/amd64,linux/arm64"

$(shell mkdir -p ${BUILD_DIR})

all: verify test build	

release:
	aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin ${RELEASE_REPO}
	BUILD_DIR=$(BUILD_DIR) RELEASE_REPO=$(RELEASE_REPO) RELEASE_PLATFORM=$(RELEASE_PLATFORM) VERSION=$(VERSION) hack/release.sh 

build:
	go build -a -ldflags="-s -w -X main.versionID=${VERSION}" -o ${BUILD_DIR}/ds-${GOOS}-${GOARCH} ${BUILD_DIR}/../cmd/ahem/main.go

test:
	go test -bench=. ${BUILD_DIR}/../... -v -coverprofile=coverage.out -covermode=atomic -outputdir=${BUILD_DIR}

verify:
	go mod tidy
	go mod download
	helm-docs
	helm-lint charts/ahem
	go vet ./...
	go fmt ./...

version:
	@echo ${VERSION}

help:
	@grep -E '^[a-zA-Z_-]+:.*$$' $(MAKEFILE_LIST) | sort


.PHONY: all build test verify help
