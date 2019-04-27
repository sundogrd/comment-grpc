package service

import (
	commentRepo "github.com/sundogrd/comment-grpc/providers/repos/comment"
	"github.com/sundogrd/comment-grpc/services/comment"
	"time"
)

type commentService struct {
	commentRepo *commentRepo.Repo
	contextTimeout  time.Duration
}

// NewUserService will create new an articleUsecase object representation of article.Usecase interface
func NewCommentService(commentRepo *commentRepo.Repo, timeout time.Duration) (comment.Service, error) {
	return &commentService{
		commentRepo: commentRepo,
		contextTimeout:  timeout,
	}, nil
}

