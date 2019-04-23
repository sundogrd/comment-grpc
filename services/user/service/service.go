package service

import (
	userRepo "github.com/sundogrd/user-grpc/providers/repos/user"
	"github.com/sundogrd/user-grpc/services/user"
	"time"
)

type userService struct {
	userRepo *userRepo.Repo
	contextTimeout  time.Duration
}

// NewUserService will create new an articleUsecase object representation of article.Usecase interface
func NewUserService(userRepo *userRepo.Repo, timeout time.Duration) (user.Service, error) {
	return &userService{
		userRepo: userRepo,
		contextTimeout:  timeout,
	}, nil
}

