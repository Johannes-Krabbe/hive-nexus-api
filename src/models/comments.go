package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comments struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt

	Content string `json:"Content" gorm:"type varchar(128) not null"`
	UserID  uint   `json:"UserID"`
	PostID  uint   `json:"PostID"`
	User    User   `json:"User"`
	Post    Post   `json:"Post"`
}
