package comment

import "time"

type CommentState int16

type Comment struct {
	// gorm.Model
	ID          int64        `gorm:"primary_key;AUTO_INCREMENT;not null"`
	AppID       string       `gorm:"type:varchar(30);index:idx_app;not null"`
	TargetID    int64        `gorm:"index:idx_target;not null;"`
	CreatorID   int64        `gorm:"not null"`
	ParentID    int64        `gorm:""`
	ReCommentID int64        `gorm:""`
	Content     string       `gorm:"type:TEXT;not null"`
	Like        int32        `gorm:"not null;DEFAULT:0"`
	Hate        int32        `gorm:"not null;DEFAULT:0"`
	Floor       int32        `gorm:"not null;DEFAULT:1"`
	State       CommentState `gorm:"type:TINYINT;NOT NULL;DEFAULT:1"`
	CreatedAt   time.Time    `gorm:"DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	ModifiedAt  time.Time    `gorm:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time   `gorm:"" sql:"index"`
	Extra       string       `gorm:"type:TEXT;"`
}

func (*Comment) TableName() string {
	return "sd_comments"
}
