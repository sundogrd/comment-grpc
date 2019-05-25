package repo

import (
	"context"
	"database/sql"
	"fmt"

	repo "github.com/sundogrd/comment-grpc/providers/repos/comment"
)

func (s commentRepo) ListRaw(ctx context.Context, req *repo.ListRawRequest) (*repo.ListResponse, error) {

	var result []*repo.Comment
	db := s.gormDB
	page := req.Page
	pageSize := req.PageSize

	var rows *sql.Rows
	var err error
	if page > 0 && pageSize > 0 {
		rows, err = db.Limit(pageSize).Offset((page-1)*pageSize).Raw(req.Query, req.Values...).Rows()
		defer rows.Close()
	} else {
		rows, err = db.Raw(req.Query, req.Values...).Rows()
		defer rows.Close()
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

	// 获取总数目
	type Total struct {
		Total int64
	}
	var total Total
	var countBeginSql string = "SELECT count(1) as total "
	var countSql = countBeginSql + req.Query[9:]

	// fmt.Println(countSql)

	db.Raw(countSql, req.Values...).Scan(&total)

	// fmt.Printf("total count is %+v", total.Total)

	res := &repo.ListResponse{
		List:     result,
		Page:     page,
		PageSize: pageSize,
		Total:    total.Total,
	}

	return res, nil

}
