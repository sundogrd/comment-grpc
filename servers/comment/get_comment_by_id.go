package comment

import (
	"context"
	"errors"
	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
)

func (server CommentServiceServer) GetCommentById(context.Context, *comment.GetCommentByIdRequest) (*comment.GetCommentByIdResponse, error) {
	return nil, errors.New("not implemented")
}