package service

import (
	"context"
	"fmt"

	"github.com/sundogrd/comment-grpc/models"
	comment "github.com/sundogrd/comment-grpc/providers/repos/comment"
	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) ListComments(ctx context.Context, req *service.ListCommentsRequest) (*service.ListCommentsResponse, error) {
	repo := *s.commentRepo

	cmts, err := repo.List(ctx, &comment.ListRequest{
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
		Sort:        comment.SortType(req.Sort),
	})
	if err != nil {
		fmt.Printf("[service/comment] GetComment: get comment error: %+v", err)
		return nil, err
	}

	var ret []*models.Comment
	for _, cmt := range cmts.List {
		ret = append(ret, &models.Comment{
			CommentId:   cmt.ID,
			TargetId:    cmt.TargetID,
			CreatorId:   cmt.CreatorID,
			ParentId:    cmt.ParentID,
			ReCommentId: cmt.ReCommentID,
			Content:     cmt.Content,
			Extra:       cmt.Extra,
			Like:        cmt.Like,
			Hate:        cmt.Hate,
			State:       int16(cmt.State),
			CreatedAt:   uint32(cmt.CreatedAt.Unix()),
			ModifiedAt:  uint32(cmt.ModifiedAt.Unix()),
			Floor:       cmt.Floor,
		})
	}

	res := &service.ListCommentsResponse{
		AppId:    req.AppId,
		Comments: ret,
		Page:     cmts.Page,
		PageSize: cmts.PageSize,
		Total:    cmts.Total,
	}
	return res, nil

}
