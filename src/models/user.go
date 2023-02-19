package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Username string `validate:"required,min=4,max=32,lowercase,alphanum" json:"Username" gorm:"type varchar(32) not null unique"`
	Email    string `validate:"required,email" json:"Email" gorm:"type varchar(128) not null unique"`
	Password string `validate:"required" json:"Password" gorm:"type varchar(128) not null"`
	Salt     string `validate:"required" json:"Salt" gorm:"type varchar(128) not null"`
}
