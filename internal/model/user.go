package model

import "github.com/google/uuid"

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
