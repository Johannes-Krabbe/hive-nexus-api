package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`

	Message string `json:"message" gorm:"type: varchar(512) not null"`
	Read    bool   `json:"read" gorm:"default: false"`

	UserID     uuid.UUID `json:"userId,omitempty"`
	ChatRoomID uuid.UUID `json:"chatRoomId,omitempty"`
	User       User      `json:"user,omitempty"`
	ChatRoom   ChatRoom  `json:"chatRoom,omitempty"`
}
