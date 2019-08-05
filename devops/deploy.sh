#!/usr/bin/env bash

# 删除原有镜像和容器
if docker ps -a | grep -q sundogrd-comment-grpc; then
    docker rm -f $(docker ps -a | grep sundogrd-comment-grpc | awk '{print $1}')
fi

if docker images | grep -q  sundogrd/comment-grpc; then
    docker rmi -f `docker images | grep sundogrd/comment-grpc | awk '{print $3}'`
fi

docker pull sundogrd/comment-grpc:$1-$2
docker run -d --name sundogrd-comment-grpc -p 24592:24592 sundogrd/comment-grpc:$1-$2
