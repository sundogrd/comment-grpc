.PHONY: init build update clean dev local-docker-build

GONAME=comment-grpc

default: build

init:
	@git submodule init && git submodule update && sh ./devops/grpc_gen.sh

update:
	@git submodule foreach git pull && sh ./devops/grpc_gen.sh

build:
	@export GO111MODULE=on && export GOPROXY=https://goproxy.cn && go build -o bin/$(GONAME)

dev:
	@export GO111MODULE=on && go run server.go

clean:
	@go clean && rm -rf ./bin/$(GONAME)

local-docker-build:
	@docker build -t comment-grpc . && docker run -d -p 8999:8999 comment-grpc