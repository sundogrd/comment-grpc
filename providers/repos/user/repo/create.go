package repo

import (
	"context"
	"fmt"
	"github.com/sundogrd/user-grpc/providers/repos/user"
)

func (s userRepo) Create(ctx context.Context, req *user.CreateRequest) (*user.CreateResponse, error) {

	userID, _ := s.idBuilder.NextId()
	userObj := user.User{
		UserID:    userID,
		Name:      req.Name,
		AvatarURL: req.AvatarURL,
		Company:   req.Company,
		Email:     req.Email,
		Extra:     req.Extra,
	}
	if dbc := s.gormDB.Create(&userObj); dbc.Error != nil {
		fmt.Printf("[services/user] Create: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}

	res := &user.CreateResponse{
		User: &user.User{
			UserID:    userObj.UserID,
			Name:      userObj.Name,
			AvatarURL: userObj.AvatarURL,
			Company:   userObj.Company,
			Email:     userObj.Email,
			CreatedAt: userObj.CreatedAt,
			UpdatedAt: userObj.UpdatedAt,
			Extra:     userObj.Extra,
		},
	}
	return res, nil
	//return nil, errors.New("not implemented")
}
