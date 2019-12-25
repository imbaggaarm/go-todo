package model

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
)

type Token struct {
	jwt.Claims
	UserID uint `json:"user_id"`
}
