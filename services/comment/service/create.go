package service

import (
	"context"
	"fmt"

	"github.com/sundogrd/comment-grpc/models"
	comment "github.com/sundogrd/comment-grpc/providers/repos/comment"
	service "github.com/sundogrd/comment-grpc/services/comment"
)

func (s *commentService) getFloor(ctx context.Context, target int64) (int32, error) {
	var values []interface{}

	repo := *s.commentRepo
	// max := 0
	cmt := comment.Comment{}
	tableName := cmt.TableName()
	sql := "SELECT * FROM "
	sql += tableName
	sql += " WHERE floor = (SELECT MAX(floor) FROM "
	sql += tableName
	sql += " WHERE target_id = ?)"

	values = append(values, target)

	listRes, err := repo.ListRaw(ctx, &comment.ListRawRequest{
		Query:  sql,
		Values: values,
	})
	if err != nil {
		return 0, err
	}

	list := listRes.List
	fmt.Printf("list is: %+v", list)

	if len(list) == 0 {
		return 0, nil
	}
	return list[0].Floor, nil
}

func (s *commentService) CreateComment(ctx context.Context, req *service.CreateRequest) (*service.CreateResponse, error) {

	repo := *s.commentRepo

	// 获取floor
	floor, floorErr := s.getFloor(ctx, req.CommentParam.TargetId)
	if floorErr != nil {
		fmt.Printf("[service/comment] CreateComment: get floor error: %+v", floorErr)
		return nil, floorErr
	}
	floor += 1
	response, err := repo.Create(ctx, &comment.CreateRequest{
		AppId: req.AppId,
		Comment: comment.CommentParams{
			TargetId:    req.CommentParam.TargetId,
			ParentId:    req.CommentParam.ParentId,
			ReCommentId: req.CommentParam.ReCommentId,
			CreatorId:   req.CommentParam.CreatorId,
			Content:     req.CommentParam.Content,
			Extra:       req.CommentParam.Extra,
			Floor:       floor,
		},
	})
	if err != nil {
		fmt.Printf("[service/comment] CreateComment: create error: %+v", err)
		return nil, err
	}

	res := &service.CreateResponse{
		AppId: response.AppId,
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
