package models

import (
	uuid "github.com/google/uuid"
)

type Output struct {
	Id          uuid.UUID `json:"id" db:"id" sql:",type:uuid"`
	Description string    `json:"description" db:"description"`
	Emails      []string  `json:"emails" db:"emails" pg:"array"`
}
