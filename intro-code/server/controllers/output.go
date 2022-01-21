package controllers

import (
	"log"

	"intro-code/intro-code/server/database"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Output struct {
	Id          uuid.UUID
	Description string   `json:"description"`
	Emails      []string `json:"emails"`
}

func CreateOutput(output Output) uuid.UUID {
	db := database.GetConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO output.output (id, description, emails) VALUES ($1, $2, $3) RETURNING id`
	output.Id = uuid.NewV4()
	err := db.QueryRow(sqlStatement, output.Id, "New Reminder sent", pq.Array(output.Emails)).Scan(&output.Id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	return output.Id
}

func GetAllOutput() []Output {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM output.output ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var outputs []Output
	for rows.Next() {
		output := Output{}

		var Uid uuid.UUID
		var description string
		var emails []string
		err := rows.Scan(&Uid, &description, pq.Array(&emails))
		if err != nil {
			log.Fatal(err)
		}

		output.Id = Uid
		output.Description = description
		output.Emails = emails

		outputs = append(outputs, output)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return outputs
}
