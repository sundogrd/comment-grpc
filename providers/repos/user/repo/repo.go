package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/sundogrd/user-grpc/providers/repos/user"
	"github.com/zheng-ji/goSnowFlake"
	"time"
)

type userRepo struct {
	gormDB *gorm.DB
	contextTimeout  time.Duration
	idBuilder *goSnowFlake.IdWorker
}

// NewUserService will create new an articleUsecase object representation of article.Usecase interface
func NewUserRepo(gormDB *gorm.DB, timeout time.Duration) (user.Repo, error) {
	idBuilder, err := goSnowFlake.NewIdWorker(3)
	if err != nil {
		return nil, err
	}
	repo := userRepo{
		gormDB: gormDB,
		contextTimeout:  timeout,
		idBuilder: idBuilder,
	}
	return repo, nil
}

