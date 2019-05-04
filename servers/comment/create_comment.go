package comment

import (
	"context"
	"fmt"

	comment "github.com/sundogrd/comment-grpc/grpc_gen/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func (server CommentServiceServer) CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (*comment.CreateCommentResponse, error) {
	service := server.CommentService

	response, err := service.CreateComment(ctx, &commentService.CreateRequest{
		AppId: req.AppId,
		CommentParam: commentService.CommentParam{
			TargetId:    req.Comment.TargetId,
			CreatorId:   req.Comment.CreatorId,
			ParentId:    req.Comment.ParentId,
			ReCommentId: req.Comment.ReCommentId,
			Content:     req.Comment.Content,
			Extra:       req.Comment.Extra,
		},
	})

	if err != nil {
		fmt.Printf("[server/comment] CreateComment: create error: %+v", err)
		return nil, err
	}

	res := &comment.CreateCommentResponse{
		AppId: response.AppId,
		Comment: &comment.Comment{
			CommentId:   response.Comment.CommentId,
			TargetId:    response.Comment.TargetId,
			CreatorId:   response.Comment.CreatorId,
			ParentId:    response.Comment.ParentId,
			ReCommentId: response.Comment.ReCommentId,
			Content:     response.Comment.Content,
			Extra:       response.Comment.Extra,
			Like:        response.Comment.Like,
			Hate:        response.Comment.Hate,
			State:       comment.Comment_ECommentState(response.Comment.State),
			CreatedAt:   response.Comment.CreatedAt,
			ModifiedAt:  response.Comment.ModifiedAt,
			Floor:       uint32(response.Comment.Floor),
		},
	}

	return res, nil
}
