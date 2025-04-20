package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `json:"id" db:"id" example:"364d2ed4-7d7e-417a-ae4e-3c59ee8dd4ee"`
	ReceptionId uuid.UUID `json:"receptionId" db:"reception_id" example:"11325f80-ef68-4176-906f-c079920953d5"`
	DateTime    time.Time `json:"dateTime" db:"date_time" example:"2025-04-20T14:43:35.824Z"`
	Type        string    `json:"type" db:"type" example:"электроника" enums:"электроника,одежда,обувь"`
}

type Products struct {
	Type  string    `json:"type" example:"электроника" enums:"электроника,одежда,обувь"`
	PvzId uuid.UUID `json:"pvzId" example:"11325f80-ef68-4176-906f-c079920953d5"`
}

type AddProduct struct {
	Type        string
	ReceptionId string
}
