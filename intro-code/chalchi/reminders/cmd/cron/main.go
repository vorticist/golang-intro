package cron

import (
	"fmt"
	"reminders/internal/repository"

	uuid "github.com/google/uuid"
)

var channelSchedule chan []uuid.UUID

func WriteToOutput(emails []uuid.UUID) {
	if len(emails) > 0 {
		fmt.Println(emails)
		channelSchedule <- emails
	}
}

func sendToOutput(emails []uuid.UUID) {
	ro := repository.NewOutputRepository()
	var output repository.Output
	output.Description = "Sent from another service"
	output.Emails = emails
	id := ro.NewOutput(output)
	fmt.Printf("new entry in output data base with id: %v\n", id)
}

func ListenChannel() {
	channelSchedule := make(chan []uuid.UUID)
	go WriteToOutput([]uuid.UUID{})
loop:
	for {
		select {
		case s1, ok := <-channelSchedule:
			if !ok {
				break loop
			}
			fmt.Println(s1)
			//Call insertar en Output
			sendToOutput(s1)
		}
	}
}
