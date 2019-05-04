package service_test

import (
	"context"
	"testing"

	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func TestCommentService_Hate(t *testing.T) {

	cs, err := initServiceTest()
	if err != nil {
		t.Fatalf("Hate Service err: %+v", err)
	}
	res, err := cs.Hate(context.Background(), &commentService.HateRequest{
		CommentId: 343181982320046080,
	})
	if err != nil {
		t.Fatalf("Hate Service err: %+v", err)
	}
	t.Logf("Hate Service: %+v", res)

}
