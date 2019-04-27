package comment

import (
	"github.com/jinzhu/gorm"
	commentRepo "github.com/sundogrd/comment-grpc/providers/repos/comment"
	commentService "github.com/sundogrd/comment-grpc/services/comment"
)

type CommentServiceServer struct{
	GormDB *gorm.DB
	CommentRepo commentRepo.Repo
	CommentService commentService.Service
}