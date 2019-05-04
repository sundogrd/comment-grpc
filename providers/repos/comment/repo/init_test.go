package repo_test

import (
	"time"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
	commentRepo "github.com/sundogrd/comment-grpc/providers/repos/comment/repo"
	"github.com/sundogrd/gopkg/db"
)

func initTestDB() (repo.Repo, error) {

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
	comment, error := commentRepo.NewCommentRepo(gormDB, 2*time.Second)
	if error != nil {
		return nil, error
	}
	return comment, nil
}
