package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=localhost port=5432 user=postgres " +
		"password=password dbname=alarmaAppGo sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}
