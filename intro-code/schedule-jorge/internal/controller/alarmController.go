package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vorticist/intro-code/schedule-jorge/alarmsApp/internal/service"
)

func AlarmController() *mux.Router {

	fmt.Println("GO running")
	router := mux.NewRouter()
	// Create alarm
	router.HandleFunc("/alarm", service.CreateAlarm).Methods("POST")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))

	return router
}
