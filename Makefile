REGISTRY      ?= 172.22.11.2:30500
VERSION       ?= shkim

PACKAGE_NAME  = github.com/cqbqdd11519/apiserver-test

SERVER_NAME  = approval-api
SERVER_IMG   = $(REGISTRY)/$(SERVER_NAME):$(VERSION)

BIN = ./build/_output/bin


.PHONY: all
all: build image push

.PHONY: build build-server
build: build-server

build-server:
	CGO_ENABLED=0 go build -o $(BIN)/approval-api $(PACKAGE_NAME)/cmd/approval-api


.PHONY: image image-server
image: image-server

image-server:
	docker build -f build/server/Dockerfile -t $(SERVER_IMG) .


.PHONY: push push-server
push: push-server

push-server:
	docker push $(SERVER_IMG)
