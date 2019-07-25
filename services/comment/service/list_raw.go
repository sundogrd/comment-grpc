package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sundogrd/comment-grpc/models"

	comment "github.com/sundogrd/comment-grpc/providers/repos/comment"
	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) ListRawComments(ctx context.Context, req *service.ListCommentsRequest) (*service.ListCommentsResponse, error) {

	var sql string = "SELECT * FROM "
	values := []interface{}{}
	cmt := comment.Comment{}
	tableName := cmt.TableName()
	sql += tableName
	sql += " WHERE "
	repo := *s.commentRepo

	sqlFragments := []string{}

	var (
		page     int32 = 0
		pageSize int32 = 0
	)

	if req.AppId != "" {
		sqlFragments = append(sqlFragments, "app_id = ? ")
		values = append(values, req.AppId)
	} else {
		fmt.Printf("[service/comment] List: must have AppId parameter")
		return nil, errors.New("app id invalid")
	}

	if req.TargetId >= 0 {
		sqlFragments = append(sqlFragments, "target_id = ? ")
		values = append(values, req.TargetId)
	}
	if req.ParentId >= 0 {
		sqlFragments = append(sqlFragments, "parent_id = ? ")
		values = append(values, req.ParentId)
	}

	if req.ReCommentId != 0 {
		sqlFragments = append(sqlFragments, "re_comment_id = ? ")
		values = append(values, req.ReCommentId)
	}

	if req.Page >= 1 {
		page = req.Page
	}
	if req.PageSize >= 1 {
		pageSize = req.PageSize
	}

	if req.CreatorId != 0 {
		sqlFragments = append(sqlFragments, "creator_id = ? ")
		values = append(values, req.CreatorId)
	}

	if req.State != 0 {
		sqlFragments = append(sqlFragments, "state = ? ")
		values = append(values, req.State)
	}

	if req.StartTime != 0 {
		sqlFragments = append(sqlFragments, "created_at > ? ")
		values = append(values, req.StartTime)
	}

	if req.EndTime != 0 {
		sqlFragments = append(sqlFragments, "created_at < ? ")
		values = append(values, req.EndTime)
	}

	// 默认不列出软删除记录
	sqlFragments = append(sqlFragments, "deleted_at is NULL")

	for idx, str := range sqlFragments {
		sql += str
		if idx != len(sqlFragments)-1 {
			sql += " AND "
		}
	}

	cmts, err := repo.ListRaw(ctx, &comment.ListRawRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Values:   values,
		Query:    sql,
	})

	if err != nil {
		fmt.Printf("[service/comment] List: list errored: %+v", err)
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
