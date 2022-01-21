package controllers

import (
	"intro-code/intro-code/server/database"
	"log"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Scheduled_items struct {
	Id          uuid.UUID
	Description string      `json:"description"`
	Users       []uuid.UUID `sql:",type:uuid[]" pg:",array"`
}

func InsertReminder(reminder Scheduled_items) uuid.UUID {
	db := database.GetConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO scheduled.scheduled_items (id, description, users) VALUES ($1, $2, $3) RETURNING id`
	reminder.Id = uuid.NewV4()
	err := db.QueryRow(sqlStatement, reminder.Id, reminder.Description, pq.Array(reminder.Users)).Scan(&reminder.Id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	return reminder.Id
}

func GetAll() []Scheduled_items {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM scheduled.scheduled_items ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var reminders []Scheduled_items
	for rows.Next() {
		reminder := Scheduled_items{}

		var Uid uuid.UUID
		var description string
		var users []uuid.UUID
		err := rows.Scan(&Uid, &description, pq.Array(&users))
		if err != nil {
			log.Fatal(err)
		}

		reminder.Id = Uid
		reminder.Description = description
		reminder.Users = users

		reminders = append(reminders, reminder)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return reminders
}
