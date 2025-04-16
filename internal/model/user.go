package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type UserClaims struct {
	jwt.StandardClaims
	Id   uuid.UUID `json:"id"`
	Role string    `json:"role"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserAnswer struct {
	Id    uuid.UUID `db:"id"`
	Email string    `db:"email"`
	Role  string    `db:"role"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginAnswer struct {
	Id       uuid.UUID `db:"id"`
	Password string    `db:"email"`
	Role     string    `db:"role"`
}
