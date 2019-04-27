package comment

import (
	"context"
)

type GetRequest struct {
	UserID int64
}
type GetResponse struct {
	Comment *Comment
}

type ListRequest struct {
	Query string
}
type ListResponse struct {
	List []*Comment
	Total int64
}

type CreateRequest struct {
	Name      string
	AvatarURL string
	Company   *string
	Email     *string
	Extra     string
}
type CreateResponse struct {
	Comment *Comment
}

type Repo interface {
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
	List(ctx context.Context, req *ListRequest) (*ListResponse, error)
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
}