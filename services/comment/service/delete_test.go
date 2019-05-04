package service_test

import (
	"context"
	"testing"

	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func TestCommentService_DeleteComment(t *testing.T) {

	cs, err := initServiceTest()
	if err != nil {
		t.Fatalf("DeleteComment Service err: %+v", err)
	}
	res, err := cs.DeleteComment(context.Background(), &commentService.DeleteRequest{
		CommentId: 343181596725096448,
	})
	if err != nil {
		t.Fatalf("DeleteComment Service err: %+v", err)
	}
	t.Logf("DeleteComment Service: %+v", res)

}
