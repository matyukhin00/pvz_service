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
	Email    string    `json:"email" db:"id"`
	Password string    `json:"password" db:"password_hash"`
	Role     string    `json:"role" db:"role"`
}

type RegisteredUser struct {
	Id    uuid.UUID
	Email string
	Role  string
}
