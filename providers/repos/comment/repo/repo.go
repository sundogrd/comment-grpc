package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/sundogrd/comment-grpc/providers/repos/comment"
	"github.com/zheng-ji/goSnowFlake"
	"time"
)

type commentRepo struct {
	gormDB *gorm.DB
	contextTimeout  time.Duration
	idBuilder *goSnowFlake.IdWorker
}

// NewUserService will create new an articleUsecase object representation of article.Usecase interface
func NewCommentRepo(gormDB *gorm.DB, timeout time.Duration) (comment.Repo, error) {
	idBuilder, err := goSnowFlake.NewIdWorker(3)
	if err != nil {
		return nil, err
	}
	repo := commentRepo{
		gormDB: gormDB,
		contextTimeout:  timeout,
		idBuilder: idBuilder,
	}
	return repo, nil
}

