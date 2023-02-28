package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"UpdatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"DeletedAt,omitempty"`

	// Title   string    `json:"Title" gorm:"type: varchar(512) not null" validate:"required,min=2,max=64"`
	Content string    `json:"Content" gorm:"type: varchar(512) not null" validate:"required,min=4,max=512"`
	UserID  uuid.UUID `json:"UserID"`
	User    User      `json:"User" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
}
