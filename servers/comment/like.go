package comment

import (
	"context"
	"fmt"

	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func (server CommentServiceServer) Like(ctx context.Context, req *comment.LikeRequest) (*comment.LikeResponse, error) {

	service := server.CommentService

	response, err := service.Like(ctx, &commentService.LikeRequest{
		CommentId: req.CommentId,
	})

	if err != nil {
		fmt.Printf("[server/comment] Like: server like error: %+v", err)
		return nil, err
	}

	res := &comment.LikeResponse{
		CommentId: response.CommentId,
	}

	return res, nil
}
