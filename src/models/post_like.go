package models

import (
	"time"

	"github.com/google/uuid"
)

type PostLike struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`

	UserID uuid.UUID `json:"UserID"`
	PostID uuid.UUID `json:"PostID"`
	User   User      `json:"User"`
	Post   Post      `json:"Post"`
}
