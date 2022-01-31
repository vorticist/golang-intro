package models

import (
	uuid "github.com/google/uuid"
	"time"
)

type Schedule struct {
	Id          uuid.UUID   `json:"id" db:"id"`
	Description string      `json:"description" db:"description"`
	Users       []uuid.UUID `json:"users" db:"users" pg:"array"`
	Date        *time.Time  `json:"date" db:"date"`
}
