package service

import (
	"context"
	"errors"
	"github.com/sundogrd/user-grpc/services/user"
)

func (s *userService) Get(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error) {

	return nil, errors.New("not implemented")
}

