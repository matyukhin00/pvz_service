package model

import (
	"time"

	"github.com/google/uuid"
)

type Reception struct {
	Id       uuid.UUID `json:"id" db:"id" example:"11325f80-ef68-4176-906f-c079920953d5"`
	DateTime time.Time `json:"dateTime" db:"date_time" example:"2025-04-20T14:32:38.032Z"`
	PvzId    uuid.UUID `json:"pvzId" db:"pvz_id" example:"fa796eea-b7f8-4426-8ea3-2884c85652fe"`
	Status   string    `json:"status" db:"status" example:"in_progress" enums:"in_progress,closed"`
}

type ReceptionInfo struct {
	Reception Reception `json:"reception"`
	Products  []Product `json:"products"`
}
