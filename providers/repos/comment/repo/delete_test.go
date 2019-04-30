package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func TestCommentProvider_Delete(t *testing.T) {
	comment, err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}

	res, err := comment.Delete(context.Background(), &repo.DeleteRequest{
		CommentId: 342857719654395904,
	})
	if err != nil {
		t.Fatalf("DeleteComment err: %+v", err)
	}
	t.Logf("DeleteComment: %+v", res)
}
