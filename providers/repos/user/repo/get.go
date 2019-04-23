package repo

import (
	"context"
	repo "github.com/sundogrd/user-grpc/providers/repos/user"
)

func (s userRepo) Get(ctx context.Context, req *repo.GetRequest) (*repo.GetResponse, error) {
	var user repo.User

	s.gormDB.Where(repo.User{
		UserID: req.UserID,
	}).First(&user)

	res := &repo.GetResponse{
		User: &user,
	}
	if res.User.UserID == 0 {
		return nil, nil
	}
	return res, nil
	//return nil, errors.New("not implemented")
}
