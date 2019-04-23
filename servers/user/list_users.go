package user

import (
	"context"
	"errors"
	"github.com/sundogrd/user-grpc/grpc_gen/user"
)

func (server UserServiceServer) ListUsers(context.Context, *user.ListUsersRequest) (*user.ListUsersResponse, error) {
	return nil, errors.New("not implemented")
}