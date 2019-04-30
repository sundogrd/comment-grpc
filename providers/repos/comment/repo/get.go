package repo

import (
	"context"
	"fmt"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func (s commentRepo) Get(ctx context.Context, req *repo.GetRequest) (*repo.GetResponse, error) {

	var comment repo.Comment
	db := s.gormDB

	dbc := db.Where(repo.Comment{
		ID: req.CommentId,
	}).First(&comment)

	if dbc.Error != nil {
		fmt.Printf("[providers/comment] Delete: db get error: %+v", dbc.Error)
		return nil, dbc.Error
	}

	res := &repo.GetResponse{
		Comment: &comment,
	}

	if res.Comment.ID == 0 {
		return nil, nil
	}

	return res, nil
}
