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
	Floor       int32
}

type GetRequest struct {
	CommentId int64
}
type GetResponse struct {
	Comment *Comment
}

type ListRawRequest struct {
	Query    string
	Page     int32
	PageSize int32
	// Receiver interface{}
	Values []interface{}
}

type SortType int16

const (
	TIME_HIGH SortType = 0
	TIME_LOW  SortType = 1
	HOT_HIGH  SortType = 2
	HOT_LOW   SortType = 3
)

type ListRequest struct {
	AppId       string
	TargetId    int64
	Page        int32
	PageSize    int32
	CreatorId   int64
	ParentId    int64
	State       int32
	StartTime   uint32
	EndTime     uint32
	ReCommentId int64
	Sort        SortType
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

type UpdateRequest struct {
	CommentId int64
	Map       map[string]interface{}
}

type UpdateResponse struct {
	Comment *Comment
}

type Repo interface {
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
	List(ctx context.Context, req *ListRequest) (*ListResponse, error)
	ListRaw(ctx context.Context, req *ListRawRequest) (*ListResponse, error)
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
	Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error)
}
