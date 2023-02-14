package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

    Username string `json:"Username" gorm:"type: varchar(32) not null unique"`
    Email string `json:"Email" gorm:"type: varchar(128) not null unique"`
	Password string `json:"Password" gorm:"type: varchar(128) not null"`
}