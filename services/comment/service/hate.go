package service

import (
	"context"
	"fmt"

	comment "github.com/sundogrd/comment-grpc/providers/repos/comment"
	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) Hate(ctx context.Context, req *service.HateRequest) (*service.HateResponse, error) {
	repo := *s.commentRepo

	cmt, err := repo.Get(ctx, &comment.GetRequest{
		CommentId: req.CommentId,
	})

	if err != nil {
		fmt.Printf("[service/comment] Hate: get error before update hate count: %+v", err)
		return nil, err
	}

	response, err := repo.Update(ctx, &comment.UpdateRequest{
		CommentId: req.CommentId,
		Map: map[string]interface{}{
			"hate": cmt.Comment.Hate + 1,
		},
	})
	if err != nil {
		fmt.Printf("[service/comment] Hate: hate error: %+v", err)
		return nil, err
	}
	res := &service.HateResponse{
		CommentId: response.Comment.ID,
	}
	return res, nil
}
