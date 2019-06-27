#!/usr/bin/env bash

git pull

docker pull breakinferno/comment-grpc:$1-$2
if docker ps -a | grep -q sundogrd-comment-grpc; then
    docker rm -f $(docker ps -a | grep sundogrd-comment-grpc | awk '{print $1}')
fi
docker run -d --name sundogrd-comment-grpc -p 24592:24592 breakinferno/comment-grpc:$1-$2
