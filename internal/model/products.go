package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `json:"id" db:"id"`
	ReceptionId uuid.UUID `json:"receptionId" db:"reception_id"`
	DateTime    time.Time `json:"dateTime" db:"date_time"`
	Type        string    `json:"type" db:"type"`
}

type AddProductInc struct {
	Type  string    `json:"type"`
	PvzId uuid.UUID `json:"pvzId"`
}

type AddProduct struct {
	Type        string
	ReceptionId string
}
