package service

import (
	"context"
	"fmt"

	comment "github.com/sundogrd/comment-grpc/providers/repos/comment"
	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) Like(ctx context.Context, req *service.LikeRequest) (*service.LikeResponse, error) {
	repo := *s.commentRepo

	cmt, err := repo.Get(ctx, &comment.GetRequest{
		CommentId: req.CommentId,
	})

	if err != nil {
		fmt.Printf("[service/comment] Like: get error before update like count: %+v", err)
		return nil, err
	}

	response, err := repo.Update(ctx, &comment.UpdateRequest{
		CommentId: req.CommentId,
		Map: map[string]interface{}{
			"like": cmt.Comment.Like + 1,
		},
	})
	if err != nil {
		fmt.Printf("[service/comment] Like: like error: %+v", err)
		return nil, err
	}
	res := &service.LikeResponse{
		CommentId: response.Comment.ID,
	}
	return res, nil
}
