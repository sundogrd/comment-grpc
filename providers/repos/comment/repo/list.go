package repo

import (
	"context"
	"fmt"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

// type CommentResult struct {
// 	Id          int64
// 	TargetId    int64
// 	CreatorId   int64
// 	ParentId    int64
// 	ReCommentId int64
// 	Content     string
// 	Extra       string
// 	Like        int32
// 	Hate        int32
// 	State       CommentState
// 	CreatedAt   time.Time
// 	ModifiedAt  time.Time
// 	Floor       uint32
// }

func (s commentRepo) List(ctx context.Context, req *repo.ListRequest) (*repo.ListResponse, error) {

	var result []*repo.Comment
	db := s.gormDB
	page := req.Page
	pageSize := req.PageSize

	rows, err := db.Raw(req.Query, req.Values...).Rows()
	defer rows.Close()

	if page > 0 && pageSize > 0 {
		db.Limit(pageSize).Offset((page - 1) * pageSize)
	}

	if err != nil {
		fmt.Printf("[providers/comment] List: db scan rows error: %+v", err)
		return nil, err
	}

	for rows.Next() {
		var commentObj repo.Comment
		if rowErr := db.ScanRows(rows, &commentObj); rowErr != nil {
			fmt.Printf("[providers/comment] List: db scan row error: %+v", rowErr)
		}
		result = append(result, &commentObj)
	}

	res := &repo.ListResponse{
		List:     result,
		Page:     page,
		PageSize: pageSize,
		Total:    int64(len(result)),
	}

	return res, nil

}
