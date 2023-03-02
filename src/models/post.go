package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`

	Title   string    `json:"title" gorm:"type: varchar(64) not null" validate:"required,min=2,max=64"`
	Content string    `json:"content" gorm:"type: varchar(512) not null" validate:"required,min=4,max=512"`
	UserID  uuid.UUID `json:"userID"`
	User    User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
}
