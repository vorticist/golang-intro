package repository

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	models "reminders/internal/models"
)

type outputRepository struct {
	db *sql.DB
}

type OutputsRepository interface {
	NewOutput(output models.Output) string
	UpdateOutput(output models.Output) int64
	DeleteOutput(id string) int64
	ListOutputs() ([]models.Output, error)
}

func (ur *outputRepository) NewOutput(output models.Output) string {
	//TODO: validate the new output not exist into data base.

	// close database
	defer ur.db.Close()

	insertStmt := `INSERT INTO ` + schemaOutput + `.` + tableOutput + ` (id, description, emails) VALUES ($1, $2, $3) RETURNING id`
	var id string

	// Scan function will save the insert id in the id
	err := ur.db.QueryRow(insertStmt, GenerateUUID(), output.Description, pq.Array(output.Emails)).Scan(&id)
	CheckError(err)
	fmt.Printf("Inserted %v  in %v\n", id, tableOutput)
	return id
}

func (ur *outputRepository) UpdateOutput(output models.Output) int64 {
	// close database
	defer ur.db.Close()

	// create the update sql query
	updateStmt := `UPDATE $1 SET description=$3, emails=$4 WHERE id=$2`

	// execute the sql statement
	res, err := ur.db.Exec(updateStmt, tableOutput, output.Id, output.Description, output.Emails)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v in %v\n", rowsAffected, tableOutput)
	return rowsAffected
}

func (ur *outputRepository) DeleteOutput(id string) int64 {
	// close database
	defer ur.db.Close()

	// create the delete sql query
	deleteStmt := `DELETE FROM $1 WHERE id=$2`
	// execute the sql statement
	res, err := ur.db.Exec(deleteStmt, tableOutput, id)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v in %v", rowsAffected, tableOutput)

	return rowsAffected
}

func (ur *outputRepository) ListOutputs() ([]models.Output, error) {
	// close database
	defer ur.db.Close()

	var outputs []models.Output

	// create the select sql query
	sqlStatement := `SELECT * FROM ` + tableOutput
	// execute the sql statement
	rows, err := ur.db.Query(sqlStatement)
	CheckError(err)
	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var output models.Output

		// unmarshal the row object to schedule
		err = rows.Scan(&output.Id, &output.Description, &output.Emails)

		CheckError(err)

		// append the output in the output slice
		outputs = append(outputs, output)
	}
	// return empty schedules on error
	return outputs, err
}

func NewOutputRepository() OutputsRepository {
	return &outputRepository{db: CreateConnection()}
}
