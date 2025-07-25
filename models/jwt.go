package models

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	UserID   uint   `json:"uid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
