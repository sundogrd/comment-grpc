package user

import (
	"context"
	"github.com/sundogrd/user-grpc/models"
)

type GetRequest struct {
	UserID int64
}
type GetResponse struct {
	UserInfo *models.UserInfo
}

//type ListRequest struct {
//
//}
//type ListResponse struct {
//	List []*models.UserInfo
//	Total int64
//}
//
//type CreateRequest struct {
//
//}
//type CreateResponse struct {
//
//}
//
//type CancelRequest struct {
//
//}
//type CancelResponse struct {
//
//}

type Service interface {
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
	//List(ctx context.Context, req ListRequest) (ListResponse, error)
	//Create(ctx context.Context, req CreateRequest) (CreateResponse, error)
	//Cancel(ctx context.Context, req CancelRequest) (CancelResponse, error)
}
