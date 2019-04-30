package service

import (
	"context"
	"fmt"

	comment "github.com/sundogrd/comment-grpc/providers/repos/comment"
	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) DeleteComment(ctx context.Context, req *service.DeleteRequest) (*service.DeleteResponse, error) {
	repo := *s.commentRepo

	response, err := repo.Delete(ctx, &comment.DeleteRequest{
		CommentId: req.CommentId,
	})
	if err != nil {
		fmt.Printf("[service/comment] DeleteComment: delete comment error: %+v", err)
		return nil, err
	}

	res := &service.DeleteResponse{
		CommentId: response.CommentId,
	}
	return res, nil
}
