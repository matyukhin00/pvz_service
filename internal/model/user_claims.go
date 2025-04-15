package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	Id   string `json:"id"`
	Role string `json:"role"`
}
