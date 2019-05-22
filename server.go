package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"time"

	commentGen "github.com/sundogrd/comment-grpc/grpc_gen/comment"
	commentRepo "github.com/sundogrd/comment-grpc/providers/repos/comment/repo"
	"github.com/sundogrd/comment-grpc/servers/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment/service"
	configUtils "github.com/sundogrd/gopkg/config"
	"github.com/sundogrd/gopkg/db"
	grpcUtils "github.com/sundogrd/gopkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := configUtils.ReadConfigFromFile("./config", nil)
	if err != nil {
		logrus.Errorf("[comment-grpc] ReadConfigFromFile err: %s", err.Error())
		panic(err)
	}

	instanceAddr := config.Get("grpcService.host").(string) + ":" + config.Get("grpcService.port").(string)
	listen, err := net.Listen("tcp", instanceAddr)
	if err != nil {
		logrus.Errorf("[comment-grpc] net.Listen err: %s", err.Error())
		panic(err)
	}

	gormDB, err := db.Connect(db.ConnectOptions{
		User:           config.Get("db.options.user").(string),
		Password:       config.Get("db.options.password").(string),
		Host:           config.Get("db.options.host").(string),
		Port:           config.Get("db.options.port").(string),
		DBName:         config.Get("db.options.dbname").(string),
		ConnectTimeout: config.Get("db.options.connectTimeout").(string),
	})
	if err != nil {
		logrus.Errorf("[comment-grpc] db.Connect err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[comment-grpc] db.Connect finished")

	cr, err := commentRepo.NewCommentRepo(gormDB, 2*time.Second)
	if err != nil {
		logrus.Errorf("[comment-grpc] NewCommentRepo err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[comment-grpc] NewCommentRepo finished")

	cs, err := commentService.NewCommentService(&cr, 2*time.Second)
	if err != nil {
		logrus.Errorf("[comment-grpc] NewCommentService err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[comment-grpc] NewCommentService finished")

	grpcServer := grpc.NewServer()
	resolver, err := grpcUtils.NewGrpcResolover()
	if err != nil {
		logrus.Errorf("[comment-grpc] NewGrpcResolover err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[comment-grpc] NewGrpcResolover finished")

	err = grpcUtils.ResgiterServer(*resolver, "sundog.comment", instanceAddr, 5*time.Second, 5)
	if err != nil {
		logrus.Errorf("[comment-grpc] RegisterServer err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[comment-grpc] ResgiterServer finished, service: %s, %s", "sundog.comment", instanceAddr)

	commentGen.RegisterCommentServiceServer(grpcServer, &comment.CommentServiceServer{
		GormDB:         gormDB,
		CommentRepo:    cr,
		CommentService: cs,
	})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
