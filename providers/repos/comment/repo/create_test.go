package repo_test

import (
	"context"
	"testing"
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

func TestCommentProvider_Create(t *testing.T) {
	comment, err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}

	res, err := comment.Create(context.Background(), &repo.CreateRequest{
		AppId: "2322",
		Comment: repo.CommentParams{
			TargetId:    232,
			CreatorId:   1111,
			ParentId:    11,
			ReCommentId: 23,
			Content:     "Test Content",
			Extra:       "Test Extra",
		},
	})
	if err != nil {
		t.Fatalf("CreateComment err: %+v", err)
	}
	t.Logf("CreateComment: %+v", res)
}
