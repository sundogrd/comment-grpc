package service_test

import (
	"context"
	"testing"

	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func TestCommentService_List(t *testing.T) {

	cs, err := initServiceTest()
	if err != nil {
		t.Fatalf("List Service err: %+v", err)
	}
	res, err := cs.ListComments(context.Background(), &commentService.ListCommentsRequest{
		AppId:     "2322",
		CreatorId: 11,
	})
	if err != nil {
		t.Fatalf("List Service err: %+v", err)
	}
	t.Logf("List Service: %+v", res)

}
