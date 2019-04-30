package service

import (
	"context"
	"errors"

	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) List(ctx context.Context, req *service.ListCommentsRequest) (*service.ListCommentsResponse, error) {

	return nil, errors.New("not implemented")
}
