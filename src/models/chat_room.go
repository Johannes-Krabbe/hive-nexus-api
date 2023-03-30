package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatRoom struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`

	Users    []User    `json:"users,omitempty" gorm:"many2many:user_chatrooms;"`
	Messages []Message `json:"messages,omitempty" gorm:"many2many:chat_room_messages;"`
}
