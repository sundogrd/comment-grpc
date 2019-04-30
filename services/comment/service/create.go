package service

import (
	"context"
	"errors"
	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) CreateComment(ctx context.Context, req *service.CreateRequest) (*service.CreateResponse, error) {

	return nil, errors.New("not implemented")
}

