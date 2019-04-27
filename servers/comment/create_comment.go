package comment

import (
	"context"
	"errors"
	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
)

func (server CommentServiceServer) CreateComment(context.Context, *comment.CreateCommentRequest) (*comment.CreateCommentResponse, error) {
	return nil, errors.New("not implemented")
}