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

type User struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password_hash"`
	Role     string    `json:"role" db:"role"`
}

type RegisteredUser struct {
	Id    uuid.UUID `json:"id" example:"ddb0897f-dfc8-4f1d-8263-f2d0d11b33fe"`
	Email string    `json:"email" example:"user@example.com"`
	Role  string    `json:"role" example:"employee" enums:"employee,moderator"`
}
