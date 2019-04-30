package comment_test

import (
	"time"

	"github.com/sundogrd/gopkg/db"

	commentRepo "github.com/sundogrd/comment-grpc/providers/repos/comment/repo"
	"github.com/sundogrd/comment-grpc/services/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment/service"
)

func initServerTest() (comment.Service, error) {

	gormDB, err := db.Connect(db.ConnectOptions{
		User:           "root",
		Password:       "12345678",
		Host:           "127.0.0.1",
		Port:           "3306",
		DBName:         "comment",
		ConnectTimeout: "10s",
	})
	if err != nil {
		return nil, err
	}

	cr, error := commentRepo.NewCommentRepo(gormDB, 2*time.Second)
	if error != nil {
		return nil, error
	}

	cs, err := commentService.NewCommentService(&cr, 2*time.Second)

	if err != nil {
		return nil, err
	}
	return cs, nil
}
