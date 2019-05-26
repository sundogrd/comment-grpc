#!/usr/bin/env bash

git pull

docker pull breakinferno/comment-grpc:$1-$2
if docker ps -a | grep -q breakinferno/comment-grpc; then
    docker rm -f $(docker ps -a | grep breakinferno/comment-grpc | awk '{print $1}')
fi
docker run -d --name sundogrd-comment-grpc -p 9431:8999 breakinferno/comment-grpc:$1-$2
