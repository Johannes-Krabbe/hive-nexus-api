package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"UpdatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"DeletedAt,omitempty"`

	Username string `validate:"required,min=4,max=32,lowercase,alphanum" json:"username" gorm:"type varchar(32) not null unique"`
	Email    string `validate:"required,email" json:"email,omitempty" gorm:"type varchar(128) not null unique"`
	Password string `validate:"required" json:"password,omitempty" gorm:"type varchar(128) not null"`
	Salt     string `validate:"required" json:"salt,omitempty" gorm:"type varchar(128) not null"`
}
