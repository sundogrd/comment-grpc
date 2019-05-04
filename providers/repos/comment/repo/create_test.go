package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func TestCommentProvider_Create(t *testing.T) {
	comment, err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}

	res, err := comment.Create(context.Background(), &repo.CreateRequest{
		AppId: "2322",
		Comment: repo.CommentParams{
			TargetId:    232,
			CreatorId:   123321,
			ParentId:    0,
			ReCommentId: 23,
			Content:     "Test Content",
			Extra:       "Test Extra",
			Floor:       1,
		},
	})
	if err != nil {
		t.Fatalf("CreateComment err: %+v", err)
	}
	t.Logf("CreateComment: %+v", res)
}
