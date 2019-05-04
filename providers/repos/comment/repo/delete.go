package repo

import (
	"context"
	"fmt"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func (s commentRepo) Delete(ctx context.Context, req *repo.DeleteRequest) (*repo.DeleteResponse, error) {
	db := s.gormDB

	dbc := db.Delete(repo.Comment{
		ID: req.CommentId,
	})

	if dbc.Error != nil {
		fmt.Printf("[providers/comment] Delete: db delete error: %+v", dbc.Error)
		return nil, dbc.Error
	}

	res := &repo.DeleteResponse{
		CommentId: req.CommentId,
	}

	return res, nil
}
