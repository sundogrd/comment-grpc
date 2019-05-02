package comment

import (
	"context"
	"errors"

	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
)

func (server CommentServiceServer) ListComments(context.Context, *comment.ListCommentsRequest) (*comment.ListCommentsResponse, error) {
	return nil, errors.New("not implemented")
}
