package api

import (
	"encoding/json"
	"intro-code/intro-code/server/controllers"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type Data_scheduled struct {
	Success bool                          `json:"success"`
	Data    []controllers.Scheduled_items `json:"data"`
	Errors  []string                      `json:"errors"`
}

type Data_users struct {
	Success bool                `json:"success"`
	Data    []controllers.Users `json:"data"`
	Errors  []string            `json:"errors"`
}

type Data_outputs struct {
	Success bool                 `json:"success"`
	Data    []controllers.Output `json:"data"`
	Errors  []string             `json:"errors"`
}

type response struct {
	ID      uuid.UUID `json:"UUID"`
	Message string    `json:"message,omitempty"`
}

func CreateReminder(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var reminder controllers.Scheduled_items

	err := json.NewDecoder(req.Body).Decode(&reminder)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := controllers.InsertReminder(reminder)

	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetReminders(w http.ResponseWriter, req *http.Request) {
	var reminder []controllers.Scheduled_items = controllers.GetAll()

	var data = Data_scheduled{true, reminder, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func CreateUser(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user controllers.Users

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := controllers.CreateUser(user)
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	var users []controllers.Users = controllers.GetAllUsers()

	var data = Data_users{true, users, make([]string, 0)}

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func CreateOutput(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var output controllers.Output

	err := json.NewDecoder(req.Body).Decode(&output)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := controllers.CreateOutput(output)
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetAllOutput(w http.ResponseWriter, req *http.Request) {
	var outputs []controllers.Output = controllers.GetAllOutput()

	var data = Data_outputs{true, outputs, make([]string, 0)}

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
