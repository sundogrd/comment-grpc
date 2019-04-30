package comment

import (
	"context"
	"fmt"

	comment "github.com/sundogrd/comment-grpc/grpc_gen/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func (server CommentServiceServer) DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest) (*comment.DeleteCommentResponse, error) {
	service := server.CommentService

	response, err := service.DeleteComment(ctx, &commentService.DeleteRequest{
		AppId:     req.AppId,
		CommentId: req.CommentId,
	})

	if err != nil {
		fmt.Printf("[server/comment] DeleteComment: delete error: %+v", err)
		return nil, err
	}

	res := &comment.DeleteCommentResponse{
		AppId:     response.AppId,
		CommentId: response.CommentId,
	}

	return res, nil
}
