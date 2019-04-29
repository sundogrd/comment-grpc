package comment

import (
	"context"
)

type CommentParams struct {
	TargetId    int64
	CreatorId   int64
	ParentId    int64
	ReCommentId int64
	Content     string
	Extra       string
}

type GetRequest struct {
	CommentId int64
}
type GetResponse struct {
	Comment *Comment
}

type ListRequest struct {
	Query string
}
type ListResponse struct {
	List     []*Comment
	Page     int32
	PageSize int32
	Total    int64
}

type CreateRequest struct {
	AppId   string
	Comment CommentParams
}
type CreateResponse struct {
	AppId   string
	Comment *Comment
}

type DeleteRequest struct {
	CommentId int64
}

type DeleteResponse struct {
	CommentId int64
}

type Repo interface {
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
	List(ctx context.Context, req *ListRequest) (*ListResponse, error)
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
}
