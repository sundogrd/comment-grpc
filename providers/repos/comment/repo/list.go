package repo

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func (s commentRepo) List(ctx context.Context, req *repo.ListRequest) (*repo.ListResponse, error) {

	db := s.gormDB

	var page int32 = 1
	var pageSize int32 = 10

	if req.Page != 0 {
		page = req.Page
	}
	if req.PageSize != 0 {
		pageSize = req.PageSize
	}

	comments := make([]*repo.Comment, 0)

	count := int64(0)
	// sort

	if req.AppId != "" {
		db = db.Where("app_id = ?", req.AppId)
	} else {
		fmt.Printf("[service/comment] List: must have AppId parameter")
		return nil, errors.New("app id invalid")
	}

	if req.TargetId >= 0 {
		fmt.Println("target_id is ", req.TargetId)
		db = db.Where("target_id = ?", req.TargetId)
	}

	if req.ParentId >= 0 {
		fmt.Println("parentId is ", req.ParentId)
		db = db.Where("parent_id = ?", req.ParentId)
	}

	if req.ReCommentId != 0 {
		db = db.Where("re_comment_id = ?", req.ReCommentId)
	}

	if req.CreatorId != 0 {
		db = db.Where("creator_id = ?", req.CreatorId)
	}

	if req.StartTime != 0 {
		db = db.Where("created_at > ?", req.StartTime)
	}

	if req.EndTime != 0 {
		db = db.Where("created_at < ?", req.EndTime)
	}

	if req.State > 0 {
		db = db.Where("state = ?", req.State)
	}

	db = db.Limit(pageSize).Offset((page - 1) * (pageSize))

	// 排序
	var sort int16 = 0
	if int16(req.Sort) != 0 {
		sort = int16(req.Sort)
	}
	switch sort {
	case 0:
		db = db.Order("created_at desc")
	case 1:
		db = db.Order("created_at")
	case 2:
		db = db.Order("like desc")
	case 3:
		db = db.Order("like")
	default:
		db = db.Order("created_at desc")
	}

	if err := db.Find(&comments).Offset(0).Limit(-1).Count(&count).Error; err != nil {
		return nil, err
	} else {
		res := &repo.ListResponse{
			List:     comments,
			Page:     page,
			PageSize: pageSize,
			Total:    count,
		}
		return res, nil
	}
}
