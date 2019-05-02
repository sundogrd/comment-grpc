package comment

import (
	"context"
	"fmt"

	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func (server CommentServiceServer) Hate(ctx context.Context, req *comment.HateRequest) (*comment.HateResponse, error) {

	service := server.CommentService

	response, err := service.Hate(ctx, &commentService.HateRequest{
		CommentId: req.CommentId,
	})

	if err != nil {
		fmt.Printf("[server/comment] Hate: server hate error: %+v", err)
		return nil, err
	}

	res := &comment.HateResponse{
		CommentId: response.CommentId,
	}

	return res, nil
}
