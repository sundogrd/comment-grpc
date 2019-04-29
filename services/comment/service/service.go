package service

import (
	"time"

	commentRepo "github.com/sundogrd/comment-grpc/providers/repos/comment"
	"github.com/sundogrd/comment-grpc/services/comment"
)

type commentService struct {
	commentRepo    *commentRepo.Repo
	contextTimeout time.Duration
}

// NewCommentService will create new an articleUsecase object representation of article.Usecase interface
func NewCommentService(commentRepo *commentRepo.Repo, timeout time.Duration) (comment.Service, error) {
	return &commentService{
		commentRepo:    commentRepo,
		contextTimeout: timeout,
	}, nil
}
