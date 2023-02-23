package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt

	Content string `json:"Content" gorm:"type varchar(128) not null"`
	UserID  uint   `json:"UserID"`
	PostID  uint   `json:"PostID"`
	User    User   `json:"User"`
	Post    Post   `json:"Post"`
}
