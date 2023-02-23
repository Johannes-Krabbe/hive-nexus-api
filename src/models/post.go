package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Content string    `json:"Content" gorm:"type: varchar(512) not null" validate:"required,min=4,max=512"`
	UserID  uuid.UUID `json:"UserID"`
	User    User      `json:"User" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
}
