package models

import "time"

type UserInfo struct {
	ID    int64      `json:"id"`
	Name      string     `json:"name"`
	AvatarURL string     `json:"avatar_url"`
	Company   *string    `json:"company"`
	Email     *string    `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Extra     string     `json:"extra"`
}