package comment

import (
	"context"
	"fmt"

	comment "github.com/sundogrd/comment-grpc/grpc_gen/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func (server CommentServiceServer) GetCommentById(ctx context.Context, req *comment.GetCommentByIdRequest) (*comment.GetCommentByIdResponse, error) {
	service := server.CommentService

	response, err := service.GetComment(ctx, &commentService.GetRequest{
		CommentId: req.CommentId,
	})

	if err != nil {
		fmt.Printf("[server/comment] GetCommentById: get by id error: %+v", err)
		return nil, err
	}

	res := &comment.GetCommentByIdResponse{
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
