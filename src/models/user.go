package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`

	Username string `validate:"required,min=4,max=32,lowercase,alphanum" json:"username" gorm:"index type varchar(32) not null unique"`
	Email    string `validate:"required,email" json:"email,omitempty" gorm:"index type varchar(128) not null unique"`
	Password string `validate:"required" json:"password,omitempty" gorm:"type varchar(128) not null"`
	Salt     string `validate:"required" json:"salt,omitempty" gorm:"type varchar(128) not null"`

	ChatRooms []ChatRoom `json:"chatRooms,omitempty" gorm:"many2many:user_chatrooms;"`
	Follows   []Follow   `json:"follows,omitempty"`
}
