package comment

import (
	"context"
	"time"

	"github.com/sundogrd/comment-grpc/models"
)

type ECommentState int32

const (
	UNKNOWN  ECommentState = 0
	SHOW     ECommentState = 1
	WITHDRAW ECommentState = 2
)

type Comment struct {
	CommentId   int64
	TargetId    int64
	CreatorId   int64
	ParentId    int64
	ReCommentId int64
	Content     string
	Extra       string
	Like        int32
	Hate        int32
	State       ECommentState
	CreatedAt   time.Time
	ModifiedAt  time.Time
	Floor       uint32
}

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
	UserInfo *models.UserInfo
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
	Comments *[]Comment
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
	Comment Comment
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
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
	// ListComments(ctx context.Context, req *ListCommentsRequest) (*ListCommentsResponse, error)
	// Like(ctx context.Context, req *LikeRequest) (*LikeResponse, error)
	// Hate(ctx context.Context, req *HateRequest) (*HateResponse, error)
	// CreateComment(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	// DeleteComment(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
}
