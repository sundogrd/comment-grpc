package comment

import (
	"context"
	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
	"testing"
)

func TestUserServiceServer_ListUsers(t *testing.T) {
	s := CommentServiceServer{}

	resp, err := s.ListComments(context.Background(), &comment.ListCommentsRequest{
		PageSize: 10,
	})
	if err != nil {
		t.Errorf("[ListCommentss] got unexpected error, %v", err)
	}
	t.Log(resp)
}