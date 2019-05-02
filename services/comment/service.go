package comment

import (
	"context"

	"github.com/sundogrd/comment-grpc/models"
)

type ECommentState int32

const (
	UNKNOWN  ECommentState = 0
	SHOW     ECommentState = 1
	WITHDRAW ECommentState = 2
)

type CommentParam struct {
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
	Comment *models.Comment
}

type ListCommentsRequest struct {
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
}

type ListCommentsResponse struct {
	AppId    string
	Comments [](*models.Comment)
	Total    int64
	Page     int32
	PageSize int32
}

type LikeRequest struct {
	CommentId int64
}

type LikeResponse struct {
	CommentId int64
}
type HateRequest struct {
	CommentId int64
}
type HateResponse struct {
	CommentId int64
}

type CreateRequest struct {
	AppId        string
	CommentParam CommentParam
}

type CreateResponse struct {
	AppId   string
	Comment *models.Comment
}

type DeleteRequest struct {
	AppId     string
	CommentId int64
}

type DeleteResponse struct {
	AppId     string
	CommentId int64
}
type Service interface {
	GetComment(ctx context.Context, req *GetRequest) (*GetResponse, error)
	ListComments(ctx context.Context, req *ListCommentsRequest) (*ListCommentsResponse, error)
	Like(ctx context.Context, req *LikeRequest) (*LikeResponse, error)
	Hate(ctx context.Context, req *HateRequest) (*HateResponse, error)
	CreateComment(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	DeleteComment(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
}
