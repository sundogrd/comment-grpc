package user

import (
	"context"
	"github.com/sundogrd/user-grpc/grpc_gen/user"
)

func (server UserServiceServer) GetUser(context.Context, *user.GetUserRequest) (*user.GetUserResponse, error) {
	//return nil, errors.New("not implemented")
	return &user.GetUserResponse{
		User: &user.User{
			Id: 123,
		},
	}, nil
}
