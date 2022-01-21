package repository

import (
	"github.com/google/uuid"
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
