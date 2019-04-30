package service

import (
	"context"
	"errors"
	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) DeleteComment(ctx context.Context, req *service.DeleteRequest) (*service.DeleteResponse, error) {

	return nil, errors.New("not implemented")
}

