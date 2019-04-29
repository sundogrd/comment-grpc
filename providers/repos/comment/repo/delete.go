package repo

import (
	"context"
	"errors"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func (s commentRepo) Delete(ctx context.Context, req *repo.DeleteRequest) (*repo.DeleteResponse, error) {

	return nil, errors.New("not implemented")
}
