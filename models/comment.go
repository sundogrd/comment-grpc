package models

type Comment struct {
	CommentId   int64  `json:"id"`
	TargetId    int64  `json:"target_id`
	CreatorId   int64  `json:"creator_id"`
	ParentId    int64  `json:"parent_id"`
	ReCommentId int64  `json:"re_comment_id"`
	Content     string `json:"content"`
	Extra       string `json:"extra"`
	Like        int32  `json:"like"`
	Hate        int32  `json:"hate"`
	State       int16  `json:"state"`
	CreatedAt   uint32 `json:"created_at"`
	ModifiedAt  uint32 `json:"modified_at"`
	Floor       int32  `json:"floor"`
}
