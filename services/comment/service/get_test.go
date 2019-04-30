package service_test

import (
	"context"
	"testing"

	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func TestCommentService_GetComment(t *testing.T) {

	cs, err := initServiceTest()
	if err != nil {
		t.Fatalf("GetComment Service err: %+v", err)
	}
	res, err := cs.GetComment(context.Background(), &commentService.GetRequest{
		CommentId: 343181596725096448,
	})
	if err != nil {
		t.Fatalf("GetComment Service err: %+v", err)
	}
	t.Logf("GetComment Service: %+v", res)

}
