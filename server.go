package main

import (
	"fmt"
	"net"
	"time"

	commentGen "github.com/sundogrd/comment-grpc/grpc_gen/comment"
	commentRepo "github.com/sundogrd/comment-grpc/providers/repos/comment/repo"
	"github.com/sundogrd/comment-grpc/servers/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment/service"
	configUtils "github.com/sundogrd/gopkg/config"
	"github.com/sundogrd/gopkg/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := configUtils.ReadConfigFromFile("./config", nil)
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", config.Get("grpcPort").(string))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	// init services and repos
	//authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	//ar := _articleRepo.NewMysqlArticleRepository(dbConn)
	//
	//timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	//au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	//_articleHttpDeliver.NewArticleHttpHandler(e, au)

	gormDB, err := db.Connect(db.ConnectOptions{
		User:           config.Get("db.options.user").(string),
		Password:       config.Get("db.options.password").(string),
		Host:           config.Get("db.options.host").(string),
		Port:           config.Get("db.options.port").(string),
		DBName:         config.Get("db.options.dbname").(string),
		ConnectTimeout: config.Get("db.options.connectTimeout").(string),
	})
	if err != nil {
		panic(err)
	}

	cr, err := commentRepo.NewCommentRepo(gormDB, 2*time.Second)
	if err != nil {
		panic(err)
	}
	cs, err := commentService.NewCommentService(&cr, 2*time.Second)

	grpcServer := grpc.NewServer()
	commentGen.RegisterCommentServiceServer(grpcServer, &comment.CommentServiceServer{
		GormDB:         gormDB,
		CommentRepo:    cr,
		CommentService: cs,
	})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
