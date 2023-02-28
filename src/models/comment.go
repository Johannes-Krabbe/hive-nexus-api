package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeleteAt  gorm.DeletedAt `json:"deletedAt,omitempty"`

	Content string    `json:"content" gorm:"type varchar(128) not null"`
	UserID  uuid.UUID `json:"userID"`
	PostID  uuid.UUID `json:"postID"`
	User    User      `json:"user"`
	Post    Post      `json:"post"`
}
