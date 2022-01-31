package repository

import (
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/lib/pq"
	models "reminders/internal/models"
)

type userRepository struct {
	db *sql.DB
}

type UsersRepository interface {
	NewUser(user models.User) string
	UpdateUser(user models.User) int64
	DeleteUser(userId string) int64
	ListUsers() ([]models.User, error)
	ExistUsers(userIds []uuid.UUID) bool
	GetEmails(userIds []uuid.UUID) []string
}

func (ur *userRepository) NewUser(user models.User) string {
	// close database
	defer ur.db.Close()
	insertStmt := `INSERT INTO ` + schemaUser + `.` + tableUser + `(user_id, email) VALUES ($1, $2) RETURNING user_id`
	var id string

	// Scan function will save the insert id in the id
	err := ur.db.QueryRow(insertStmt, user.IdUser, user.Email).Scan(&id)
	CheckError(err)
	fmt.Printf("Inserted %v in %v\n", id, tableUser)
	return id
}

func (ur *userRepository) UpdateUser(user models.User) int64 {
	// close database
	defer ur.db.Close()

	// create the update sql query
	updateStmt := `UPDATE ` + schemaUser + `.` + tableUser + ` SET email=$2  WHERE user_id=$1`

	// execute the sql statement
	res, err := ur.db.Exec(updateStmt, user.IdUser, user.Email)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}

func (ur *userRepository) DeleteUser(userId string) int64 {
	// create the postgres db connection
	//ur.db = CreateConnection(dbnameUser)

	// close database
	defer ur.db.Close()

	// create the delete sql query
	deleteStmt := `DELETE FROM ` + schemaUser + `.` + tableUser + ` WHERE user_id=$1`
	// execute the sql statement
	res, err := ur.db.Exec(deleteStmt, userId)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func (ur *userRepository) ListUsers() ([]models.User, error) {
	// create the postgres db connection
	//ur.db = CreateConnection(dbnameUser)

	// close database
	defer ur.db.Close()

	var users []models.User

	// create the select sql query
	sqlStatement := `SELECT * FROM ` + schemaUser + `.` + tableUser
	// execute the sql statement
	rows, err := ur.db.Query(sqlStatement)
	CheckError(err)
	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user models.User

		// unmarshal the row object to user
		err = rows.Scan(&user.IdUser, &user.Email)

		CheckError(err)
		// append the user in the users slice
		users = append(users, user)
	}
	// return empty users on error
	return users, err
}

func (ur *userRepository) ExistUsers(userIds []uuid.UUID) bool {
	var result bool = false
	// close database
	defer ur.db.Close()
	queryStmt := `SELECT EXISTS(SELECT 1 FROM ` + schemaUser + `.` + tableUser + ` WHERE user_id::text = ANY ($1) HAVING COUNT (user_id) = $2);`	
	res, err := ur.db.Query(queryStmt, pq.Array(userIds), len(userIds))
	CheckError(err)
	res.Scan(&result)	
	return result
}

func (ur *userRepository) GetEmails(userIds []uuid.UUID) []string {
	// close database
	defer ur.db.Close()

	var emails []string
	// create the select sql query
	sqlStatement := `SELECT email FROM ` + schemaUser + `.` + tableUser + ` WHERE  user_id::text = ANY ($1)`
	// execute the sql statement
	rows, err := ur.db.Query(sqlStatement, pq.Array(userIds))
	CheckError(err)
	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var email string

		// unmarshal the row object to user
		err = rows.Scan(&email)

		CheckError(err)
		// append the user in the users slice
		emails = append(emails, email)
	}
	// return empty emails on error
	return emails
}

func NewUserRepository() UsersRepository {
	return &userRepository{db: CreateConnection()}
}
