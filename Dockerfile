FROM golang:latest
MAINTAINER Breakinferno <1972952841@qq.com>

COPY . $GOPATH/src/github.com/sundogrd/comment-grpc
WORKDIR $GOPATH/src/github.com/sundogrd/comment-grpc

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor


RUN go build .

RUN "./devops/build_docker.sh"

EXPOSE 24952
  
ENTRYPOINT ["./devops/entrypoint.sh"]