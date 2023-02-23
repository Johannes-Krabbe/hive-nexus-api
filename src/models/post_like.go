package models

import (
	"time"

	"gorm.io/gorm"
)

type PostLike struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt

	UserID uint `json:"UserID"`
	PostID uint `json:"PostID"`
	User   User `json:"User"`
	Post   Post `json:"Post"`
}
