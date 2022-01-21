package controllers

import (
	"intro-code/intro-code/server/database"
	"log"

	uuid "github.com/satori/go.uuid"
)

type Users struct {
	Id    uuid.UUID
	Email string `json:"email"`
}

func CreateUser(user Users) uuid.UUID {
	db := database.GetConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO public.users (user_id, email) VALUES ($1, $2) RETURNING user_id`
	user.Id = uuid.NewV4()
	err := db.QueryRow(sqlStatement, user.Id, user.Email).Scan(&user.Id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	return user.Id
}

func GetAllUsers() []Users {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM public.users ORDER BY user_id")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var users []Users
	for rows.Next() {
		user := Users{}

		var Uid uuid.UUID
		var email string
		err := rows.Scan(&Uid, &email)
		if err != nil {
			log.Fatal(err)
		}

		user.Id = Uid
		user.Email = email

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return users
}
