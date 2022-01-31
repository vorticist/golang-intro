package repository

import (
	"github.com/google/uuid"
)

const (
	schemaUser     = "users"
	tableUser      = "users"
	schemaSchedule = "schedules"
	tableSchedule  = "scheduled_items"
	schemaOutput   = "outputs"
	tableOutput    = "output"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateUUID() uuid.UUID {
	uuidNew := uuid.New()
	return uuidNew
}
