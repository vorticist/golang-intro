package models

import (
	uuid "github.com/google/uuid"
)


type User struct {
	IdUser uuid.UUID `json:"user_id" db:"user_id" sql:",type:uuid"`
	Email  string    `json:"email" db:"email"`
}
