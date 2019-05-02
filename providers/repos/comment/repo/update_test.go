package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func TestCommentProvider_Update(t *testing.T) {
	comment, err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}

	res, err := comment.Update(context.Background(), &repo.UpdateRequest{
		CommentId: 343191254370103296,
		Map: map[string]interface{}{
			"like": 3,
			"hate": 2,
		},
	})
	if err != nil {
		t.Fatalf("ListComment err: %+v", err)
	}

	t.Logf("ListComment: %+v", res)
}
