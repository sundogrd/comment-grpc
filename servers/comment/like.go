package comment

import (
	"context"
	"errors"
	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
)

func (server CommentServiceServer) Like(context.Context, *comment.LikeRequest) (*comment.LikeResponse, error) {
	return nil, errors.New("not implemented")
}