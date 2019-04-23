package user

import (
	"github.com/jinzhu/gorm"
	userRepo "github.com/sundogrd/user-grpc/providers/repos/user"
	userService "github.com/sundogrd/user-grpc/services/user"
)

type UserServiceServer struct{
	GormDB *gorm.DB
	UserRepo userRepo.Repo
	UserService userService.Service
}