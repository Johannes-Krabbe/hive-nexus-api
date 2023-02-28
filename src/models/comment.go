package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"UpdatedAt,omitempty"`
	DeleteAt  gorm.DeletedAt `json:"DeletedAt,omitempty"`

	Content string    `json:"Content" gorm:"type varchar(128) not null"`
	UserID  uuid.UUID `json:"UserID"`
	PostID  uuid.UUID `json:"PostID"`
	User    User      `json:"User"`
	Post    Post      `json:"Post"`
}
