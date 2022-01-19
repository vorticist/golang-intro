package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)


type User struct {
	id   string
	email string
}

func save_output(emails []string) {
	uri := "postgres://root:toor@127.0.0.1:5432/go?sslmode=disable"
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	sqlStatement := "INSERT INTO output (description, emails) VALUES ($1, $2)"
	_, err = db.Exec(sqlStatement, "Alarma", pq.Array(emails))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OUTPUT Saved")
}

func save_schedule(c echo.Context) error{

	var uuid_context []string

	err := echo.QueryParamsBinder(c).
		Strings("uuid", &uuid_context).
		BindError()

	if err != nil {
		return c.String(http.StatusBadRequest, "Error")
	}

	var emails []string
	for _, element := range uuid_context{
		urls := fmt.Sprintf("http://localhost:123/users?uuid=%s", element)
		resp, err := http.Get(urls)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error")
		}
		if resp.StatusCode != 200{
			return c.String(http.StatusNoContent, "Not found")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.String(http.StatusBadRequest, "Error")
		}
		sb := string(body)
		emails = append(emails, sb)
	}

	uri := "postgres://root:toor@127.0.0.1:5432/go?sslmode=disable"
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error")
	}
	sqlStatement := "INSERT INTO scheduled_items (description, users) VALUES ($1, $2)"
	_, err = db.Exec(sqlStatement, "Alarma", pq.Array(uuid_context))
	if err != nil {
		return c.String(http.StatusBadRequest, "Error")
	}

	time.AfterFunc(10*time.Second, func(){save_output(emails)})

	return c.String(http.StatusOK, "Done, message set")
}

func main() {
	e := echo.New()
	e.POST("/schedule", save_schedule)
	e.Logger.Fatal(e.Start(":1234"))
}
