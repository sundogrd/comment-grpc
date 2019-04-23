package user

import (
	"context"
	"github.com/sundogrd/user-grpc/grpc_gen/user"
	"testing"
)

func TestUserServiceServer_ListUsers(t *testing.T) {
	s := UserServiceServer{}

	resp, err := s.ListUsers(context.Background(), &user.ListUsersRequest{
		PageSize: 10,
	})
	if err != nil {
		t.Errorf("[ListUsers] got unexpected error, %v", err)
	}
	t.Log(resp)
}