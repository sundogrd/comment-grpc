package repo

import (
	"context"
	"fmt"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func (s commentRepo) Update(ctx context.Context, req *repo.UpdateRequest) (*repo.UpdateResponse, error) {

	db := s.gormDB

	dbc := db.Model(&repo.Comment{}).Where("id = ?", req.CommentId).Update(req.Map)

	if dbc.Error != nil {
		fmt.Printf("[providers/comment] Update: db update error: %+v", dbc.Error)
		return nil, dbc.Error
	}

	response, err := s.Get(ctx, &repo.GetRequest{
		CommentId: req.CommentId,
	})

	if err != nil {
		fmt.Printf("[providers/comment] Update: db get comment after update error: %+v", err)
		return nil, err
	}

	res := &repo.UpdateResponse{
		Comment: response.Comment,
	}

	return res, nil

}
