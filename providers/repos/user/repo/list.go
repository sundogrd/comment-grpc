package repo

import (
	"context"
	"errors"
	"github.com/sundogrd/user-grpc/providers/repos/user"
)

func (s userRepo) List(ctx context.Context, req *user.ListRequest) (*user.ListResponse, error) {

	return nil, errors.New("not implemented")
}
