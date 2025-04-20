package model

import "github.com/google/uuid"

type Register struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password"`
	Role     string `json:"role" example:"employee" enums:"employee,moderator"`
}

type DummyLogin struct {
	Role string `json:"role" example:"employee" enums:"employee,moderator"`
}

type Login struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password"`
}

type Receptions struct {
	Id uuid.UUID `json:"pvzId" example:"fa796eea-b7f8-4426-8ea3-2884c85652fe"`
}
