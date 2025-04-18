package model

import (
	"time"

	"github.com/google/uuid"
)

type Reception struct {
	Id       uuid.UUID `json:"id" db:"id"`
	DateTime time.Time `json:"dateTime" db:"date_time"`
	PvzId    uuid.UUID `json:"pvzId" db:"pvz_id"`
	Status   string    `json:"status" db:"status"`
}

type ReceptionInfo struct {
	Reception Reception `json:"reception"`
	Products  []Product `json:"products"`
}
