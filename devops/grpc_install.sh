#!/usr/bin/env bash

# 安装grpc工具protoc,不然不能生成protobuf

PROTOBUF_VERSION=3.6.0
PROTOC_FILENAME=protoc-${PROTOBUF_VERSION}-linux-x86_64.zip

cd /home/travis
wget https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/${PROTOC_FILENAME}
unzip ${PROTOC_FILENAME}
bin/protoc --version
popd