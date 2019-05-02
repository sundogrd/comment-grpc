package service_test

import (
	"context"
	"testing"

	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func TestCommentService_Like(t *testing.T) {

	cs, err := initServiceTest()
	if err != nil {
		t.Fatalf("Like Service err: %+v", err)
	}
	res, err := cs.Like(context.Background(), &commentService.LikeRequest{
		CommentId: 343181982320046080,
	})
	if err != nil {
		t.Fatalf("Like Service err: %+v", err)
	}
	t.Logf("Like Service: %+v", res)

}
