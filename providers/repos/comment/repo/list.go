package repo

import (
	"context"
	"errors"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func (s commentRepo) List(ctx context.Context, req *repo.ListRequest) (*repo.ListResponse, error) {
	// db := s.gormDB
	// db.R(req.Query)
	return nil, errors.New("not implemented")
}
