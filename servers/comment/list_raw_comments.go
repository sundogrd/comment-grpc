package comment

import (
	"context"
	"fmt"

	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

func (server CommentServiceServer) ListRawComments(ctx context.Context, req *comment.ListCommentsRequest) (*comment.ListCommentsResponse, error) {

	service := server.CommentService

	response, err := service.ListRawComments(ctx, &commentService.ListCommentsRequest{
		AppId:       req.AppId,
		TargetId:    req.TargetId,
		Page:        req.Page,
		PageSize:    req.PageSize,
		CreatorId:   req.CreatorId,
		ParentId:    req.ParentId,
		State:       req.State,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		ReCommentId: req.ReCommentId,
		Sort:        int16(req.Sort),
	})

	if err != nil {
		fmt.Printf("[server/comment] List: server list error: %+v", err)
		return nil, err
	}

	var ret []*comment.Comment
	for _, cmt := range response.Comments {
		ret = append(ret, &comment.Comment{
			CommentId:   cmt.CommentId,
			TargetId:    cmt.TargetId,
			CreatorId:   cmt.CreatorId,
			ParentId:    cmt.ParentId,
			ReCommentId: cmt.ReCommentId,
			Content:     cmt.Content,
			Extra:       cmt.Extra,
			Like:        cmt.Like,
			Hate:        cmt.Hate,
			State:       comment.Comment_ECommentState(cmt.State),
			CreatedAt:   cmt.CreatedAt,
			ModifiedAt:  cmt.ModifiedAt,
			Floor:       uint32(cmt.Floor),
		})
	}

	res := &comment.ListCommentsResponse{
		AppId:    response.AppId,
		Total:    response.Total,
		Page:     response.Page,
		PageSize: response.PageSize,
		Comments: ret,
	}

	return res, nil
}
