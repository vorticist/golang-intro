package main

import (
    "log"
	"net/http"
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)


type User struct {
	id   string
	email string
}

func get_valid_uuid(c echo.Context) error{
	uuid_context := c.FormValue("uuid")

	if len(uuid_context) > 0 {
		uri := "postgres://root:toor@127.0.0.1:5432/go?sslmode=disable"
		db, err := sql.Open("postgres", uri)
		if err != nil {
			log.Fatal(err)
		}

		var user User
		queryErr := db.QueryRow("SELECT * FROM users WHERE user_id = $1", uuid_context).Scan(&user.id, &user.email)
		if queryErr != nil{
			fmt.Println(queryErr)
			return c.String(http.StatusNoContent, "Not found")
		}
		return c.String(http.StatusOK, string(user.email))
	}
	return c.String(http.StatusBadRequest, "Error")
}

func main() {
	e := echo.New()
	e.GET("/users", get_valid_uuid)
	e.Logger.Fatal(e.Start(":123"))
}
