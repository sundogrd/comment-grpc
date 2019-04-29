package repo

import (
	"context"
	"fmt"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

// 更新获取已有评论楼层
func (s commentRepo) getMaxFloor(appId string, targetId int64) (int32, error) {
	// 这里先从数据库里面进行查询吧，可改为查redis

	return 1, nil
}

func (s commentRepo) Create(ctx context.Context, req *repo.CreateRequest) (*repo.CreateResponse, error) {

	var floor int32
	var err error
	db := s.gormDB
	commentId, _ := s.idBuilder.NextId()
	commentParams := req.Comment
	// 判断是否是评论，从而获取层数
	if commentParams.ParentId == 0 {
		floor = 0
	} else {
		floor, err = s.getMaxFloor(req.AppId, commentParams.TargetId)
		if err != nil {
			fmt.Printf("[providers/comment] getMaxFloor: db getMaxFloor error: %+v", err)
			return nil, err
		}
	}
	commentObj := repo.Comment{
		ID:          commentId,
		AppID:       req.AppId,
		TargetID:    commentParams.TargetId,
		CreatorID:   commentParams.CreatorId,
		ParentID:    commentParams.ParentId,
		ReCommentID: commentParams.ReCommentId,
		Content:     commentParams.Content,
		Extra:       commentParams.Extra,
		Floor:       floor,
	}
	if dbc := db.Create(&commentObj); dbc.Error != nil {
		fmt.Printf("[providers/comment] Create: db createerror: %+v", dbc.Error)
		return nil, dbc.Error
	}

	res := &repo.CreateResponse{
		AppId: req.AppId,
		Comment: &repo.Comment{
			ID:          commentObj.ID,
			AppID:       commentObj.AppID,
			TargetID:    commentObj.TargetID,
			CreatorID:   commentObj.CreatorID,
			ParentID:    commentObj.ParentID,
			ReCommentID: commentObj.ReCommentID,
			Content:     commentObj.Content,
			Extra:       commentObj.Extra,
			Floor:       commentObj.Floor,
			CreatedAt:   commentObj.CreatedAt,
			ModifiedAt:  commentObj.ModifiedAt,
			DeletedAt:   commentObj.DeletedAt,
			State:       commentObj.State,
			Hate:        commentObj.Hate,
			Like:        commentObj.Like,
		},
	}
	return res, nil
}
