package model

import (
	"time"

	"github.com/google/uuid"
)

type Pvz struct {
	Id               uuid.UUID `json:"id" db:"id" example:"11325f80-ef68-4176-906f-c079920953d5"`
	RegistrationDate time.Time `json:"registrationDate" db:"registration_date" example:"2025-04-20T14:26:22.671Z"`
	City             string    `json:"city" db:"city" example:"Москва" enums:"Москва,Казань,Санкт-Петербург"`
}

type PvzInfo struct {
	Pvz        Pvz             `json:"pvz"`
	Receptions []ReceptionInfo `json:"receptions"`
}
