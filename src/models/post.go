package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

    Content string `json:"Content" gorm:"type: varchar(512) not null" `
    UserID uint `json:"UserID"`
    User User `json:"User" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
}