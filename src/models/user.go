package models

import "gorm.io/gorm"

type User struct {
    gorm.Model  // adds ID, created_at etc.
    Username		string `json:"username"`
    Email	 	string `json:"email"`
	Password	string `json:"password"`
}