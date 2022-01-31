package repository

import (
	"database/sql"
	"fmt"
	models "reminders/internal/models"
	"time"

	"github.com/lib/pq"
)

type scheduleRepository struct {
	db *sql.DB
}

type SchedulesRepository interface {
	NewSchedule(schedule models.Schedule) string
	UpdateSchedule(schedule models.Schedule) int64
	DeleteSchedule(id string) int64
	ListSchedules() ([]models.Schedule, error)
}

func (ur *scheduleRepository) NewSchedule(schedule models.Schedule) string {
	//TODO: validate the new schedule not exist into data base.
	// close database
	defer ur.db.Close()

	insertStmt := `INSERT INTO ` + schemaSchedule + `.` + tableSchedule + ` (id, description, users, date) VALUES ($1, $2, $3, $4) RETURNING id`
	var id string
	// Scan function will save the insert id in the id
	err := ur.db.QueryRow(insertStmt, schedule.Id, schedule.Description, pq.Array(schedule.Users), time.Now()).Scan(&id)
	CheckError(err)
	fmt.Printf("Inserted %v in %v\n", id, tableSchedule)
	return id
}

func (ur *scheduleRepository) UpdateSchedule(schedule models.Schedule) int64 {
	// close database
	defer ur.db.Close()

	// create the update sql query
	updateStmt := `UPDATE ` + schemaSchedule + `.` + tableSchedule + ` SET description=$2, users=$3 WHERE id=$1`

	// execute the sql statement
	res, err := ur.db.Exec(updateStmt, schedule.Id, schedule.Description, schedule.Users)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v in %v\n", rowsAffected, tableSchedule)
	return rowsAffected
}

func (ur *scheduleRepository) DeleteSchedule(id string) int64 {
	// close database
	defer ur.db.Close()

	// create the delete sql query
	deleteStmt := `DELETE FROM ` + schemaSchedule + `.` + tableSchedule + ` WHERE id=$1`
	// execute the sql statement
	res, err := ur.db.Exec(deleteStmt, id)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func (ur *scheduleRepository) ListSchedules() ([]models.Schedule, error) {
	// close database
	defer ur.db.Close()

	var schedules []models.Schedule

	// create the select sql query
	sqlStatement := `SELECT * FROM ` + schemaSchedule + `.` + tableSchedule
	// execute the sql statement
	rows, err := ur.db.Query(sqlStatement)
	CheckError(err)
	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var schedule models.Schedule

		// unmarshal the row object to schedule
		err = rows.Scan(&schedule.Id, &schedule.Description, &schedule.Users)

		CheckError(err)

		// append the schedule in the schedules slice
		schedules = append(schedules, schedule)
	}
	// return empty schedules on error
	return schedules, err
}

func NewScheduleRepository() SchedulesRepository {
	return &scheduleRepository{db: CreateConnection()}
}
