package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"-"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RequestToUser(in AuthRequest) User {
	user := User{
		Username: in.Username,
		Password: in.Password,
	}
	return user
}
