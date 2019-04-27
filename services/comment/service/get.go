package service

import (
	"context"
	"errors"
	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) Get(ctx context.Context, req *service.GetRequest) (*service.GetResponse, error) {

	return nil, errors.New("not implemented")
}

