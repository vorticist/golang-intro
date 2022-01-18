package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vorticist/intro-code/schedule-jorge/alarmsApp/internal/database"
)

func getConnection() {
	db := database.GetDB()
	fmt.Println(db)
}

func createAlarm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create alarm")
}

func main() {
	getConnection()

	router := mux.NewRouter()
	// Read
	router.HandleFunc("/alarm", createAlarm).Methods("POST")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))

}
