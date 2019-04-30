package service

import (
	"context"
	"fmt"

	"github.com/sundogrd/comment-grpc/models"
	comment "github.com/sundogrd/comment-grpc/providers/repos/comment"
	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) GetComment(ctx context.Context, req *service.GetRequest) (*service.GetResponse, error) {
	repo := *s.commentRepo

	response, err := repo.Get(ctx, &comment.GetRequest{
		CommentId: req.CommentId,
	})
	if err != nil {
		fmt.Printf("[service/comment] GetComment: get comment error: %+v", err)
		return nil, err
	}
	res := &service.GetResponse{
		Comment: &models.Comment{
			CommentId:   response.Comment.ID,
			TargetId:    response.Comment.TargetID,
			CreatorId:   response.Comment.CreatorID,
			ParentId:    response.Comment.ParentID,
			ReCommentId: response.Comment.ReCommentID,
			Content:     response.Comment.Content,
			Extra:       response.Comment.Extra,
			Like:        response.Comment.Like,
			Hate:        response.Comment.Hate,
			State:       int16(response.Comment.State),
			CreatedAt:   uint32(response.Comment.CreatedAt.Unix()),
			ModifiedAt:  uint32(response.Comment.ModifiedAt.Unix()),
			Floor:       response.Comment.Floor,
		},
	}
	return res, nil
}
