package comment_test

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/sundogrd/gopkg/db"

	commentGen "github.com/sundogrd/comment-grpc/grpc_gen/comment"

	commentRepo "github.com/sundogrd/comment-grpc/providers/repos/comment/repo"
	"github.com/sundogrd/comment-grpc/servers/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address     = "localhost:8999"
	defaultName = "world"
)

func initServer(message chan bool) {
	listen, err := net.Listen("tcp", ":8999")
	fmt.Println("开始启动服务器......")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
	}

	gormDB, err := db.Connect(db.ConnectOptions{
		User:           "root",
		Password:       "12345678",
		Host:           "127.0.0.1",
		Port:           "3306",
		DBName:         "comment",
		ConnectTimeout: "10s",
	})
	if err != nil {
		log.Fatalf("start test server failed! %+v", err)
	}

	cr, error := commentRepo.NewCommentRepo(gormDB, 2*time.Second)
	if error != nil {
		log.Fatalf("start test server failed! %+v", error)
	}

	cs, err := commentService.NewCommentService(&cr, 2*time.Second)

	if err != nil {
		log.Fatalf("start test server failed! %+v", err)
	}

	grpcServer := grpc.NewServer()
	commentGen.RegisterCommentServiceServer(grpcServer, &comment.CommentServiceServer{
		GormDB:         gormDB,
		CommentRepo:    cr,
		CommentService: cs,
	})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
	message <- true
}
