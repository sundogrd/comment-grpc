package comment

import (
	"context"
	"errors"
	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
)

func (server CommentServiceServer) Hate(context.Context, *comment.HateRequest) (*comment.HateResponse, error) {
	return nil, errors.New("not implemented")
}