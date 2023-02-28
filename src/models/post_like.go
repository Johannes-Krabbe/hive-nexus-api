package models

import (
	"time"

	"github.com/google/uuid"
)

type PostLike struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	UserID uuid.UUID `json:"userID"`
	PostID uuid.UUID `json:"postID"`
	User   User      `json:"user"`
	Post   Post      `json:"post"`
}
