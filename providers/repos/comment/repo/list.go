package repo

import (
	"context"
	"errors"
	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func (s commentRepo) List(ctx context.Context, req *repo.ListRequest) (*repo.ListResponse, error) {

	return nil, errors.New("not implemented")
}
