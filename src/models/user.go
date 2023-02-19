package models

import (
	"time"

	"gorm.io/gorm"

	valid "github.com/asaskevich/govalidator"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Username string `json:"Username" gorm:"type: varchar(32) not null unique valid:"usernameValidator, length(4|32)"`
	Email    string `json:"Email" gorm:"type: varchar(128) not null unique" valid:"email"`
	Password string `json:"Password" gorm:"type: varchar(128) not null"`
	Salt     string `json:"Salt" gorm:"type: varchar(128) not null"`
}

func init() {
	valid.CustomTypeTagMap.Set("usernameValidator", func(i interface{}, context interface{}) bool {
		switch v := context.(type) {
		case User:
			for _, c := range v.Username {
				if valid.IsNumeric(string(c)) || valid.IsLowerCase(string(c)) || string(c) == "-" || string(c) == "_" {
					return false
				}
			}
			return true
		}
		return false
	})
}
