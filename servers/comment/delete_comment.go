package comment

import (
	"context"
	"errors"
	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
)

func (server CommentServiceServer) DeleteComment(context.Context, *comment.DeleteCommentRequest) (*comment.DeleteCommentResponse, error) {
	return nil, errors.New("not implemented")
}