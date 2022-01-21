package vo

import (
	uuid "github.com/google/uuid"
)

type User struct {
	IdUser uuid.UUID `json:"userid"`
	Email  string    `json:"email"`
}

type Schedule struct {
	Id          uuid.UUID   `json:"id"`
	Description string      `json:"description"`
	Users       []uuid.UUID `json:"users" pg:"array"`
}

type Output struct {
	Id          uuid.UUID   `json:"id"`
	Description string      `json:"description"`
	Emails      []uuid.UUID `json:"emails" pg:"array"`
}
