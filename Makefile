GONAME=music-grpc

default: build

init:
	@sh ./devops/grpc_gen.sh

update:
	@git submodule foreach git pull

start:


dev:
	@go run server.go
