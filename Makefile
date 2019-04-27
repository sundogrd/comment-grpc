GONAME=music-grpc

default: build

init:
	@sh ./devops/grpc_gen.sh

update:
	@git submodule foreach git pull && sh ./devops/grpc_gen.sh

start:


dev:
	@go run server.go
