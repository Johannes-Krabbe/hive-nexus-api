package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Follow struct {
	ID        uuid.UUID      `gorm:"primaryKey;type uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeleteAt  gorm.DeletedAt `json:"deletedAt,omitempty"`

	Followee   *User     `json:"followee,omitempty" gorm:"type not null"`
	FolloweeID uuid.UUID `json:"followeeId,omitempty" gorm:"type not null"`
	Accepted   bool      `json:"accepted,omitempty" gorm:"default:false"`

	UserID uuid.UUID `json:"userId,omitempty"`
	User   User      `json:"user,omitempty"`
}
