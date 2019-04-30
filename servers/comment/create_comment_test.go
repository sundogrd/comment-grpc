package comment_test

import (
	"context"
	"testing"

	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func TestCommentServer_Create(t *testing.T) {

	cs, err := initServerTest()
	if err != nil {
		t.Fatalf("CreateComment Service err: %+v", err)
	}
	res, err := cs.CreateComment(context.Background(), &commentService.CreateRequest{
		AppId: "service test",
		CommentParam: commentService.CommentParam{
			TargetId:    1111,
			CreatorId:   1112,
			ParentId:    0,
			ReCommentId: 0,
			Content:     "Test Service Content",
			Extra:       "Test Service Extra",
		},
	})
	if err != nil {
		t.Fatalf("CreateComment Service err: %+v", err)
	}
	t.Logf("CreateComment Service: %+v", res)

}
