package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func TestCommentProvider_List(t *testing.T) {
	comment, err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}

	var values []interface{}

	values = append(values, 232)
	res, err := comment.List(context.Background(), &repo.ListRequest{
		Query:  "SELECT * FROM sd_comments WHERE target_id = ?",
		Values: values,
	})
	if err != nil {
		t.Fatalf("ListComment err: %+v", err)
	}

	t.Logf("ListComment: %+v", res)
}
