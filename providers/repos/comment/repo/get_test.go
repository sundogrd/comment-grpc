package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func TestCommentProvider_Get(t *testing.T) {
	comment, err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}

	res, err := comment.Get(context.Background(), &repo.GetRequest{
		CommentId: 342857719654395904,
	})
	if err != nil {
		t.Fatalf("GetComment err: %+v", err)
	}
	t.Logf("GetComment: %+v", res)
}
