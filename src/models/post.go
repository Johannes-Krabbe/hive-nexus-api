package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Content string `json:"Content" gorm:"type: varchar(512) not null" validate:"required,min=4,max=512"`
	UserID  uint   `json:"UserID"`
	User    User   `json:"User" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
}
