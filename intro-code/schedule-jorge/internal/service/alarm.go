package service

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/vorticist/intro-code/schedule-jorge/alarmsApp/internal/database"
)

type Users struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type Scheduled_items struct {
	UserUuid    string   `json:"userUuid"`
	AlarmAt     string   `json:"alarmAt"`
	Description string   `json:"description"`
	Uuids       []string `json:"uuids"`
	AlarmId     string   `json:"alarmId"`
}

func getConnection() *sql.DB {
	return database.GetDB()
}

func CreateAlarm(w http.ResponseWriter, r *http.Request) {
	db := getConnection()
	defer db.Close()

	var user Users
	var users []Users
	var scheduled_items Scheduled_items
	json.NewDecoder(r.Body).Decode(&scheduled_items)

	alarmAtTime, _ := time.Parse("2006-01-02 15:04:05", scheduled_items.AlarmAt)
	strNow := time.Now().Format("2006-01-02 15:04:05")
	timeNow, _ := time.Parse("2006-01-02 15:04:05", strNow)

	if alarmAtTime.Sub(timeNow).Seconds() > 0 {

		row := db.QueryRow("SELECT * FROM users.users WHERE user_id = '" + scheduled_items.UserUuid + "'")
		err := row.Scan(&user.Id, &user.Email)
		if err != nil {
			errorManager(w, "No se encontró el usuario")
			// panic(err)
		} else {

			var tempItems string
			for _, item := range scheduled_items.Uuids {
				tempItems = tempItems + "'" + item + "',"
			}
			tempItems = tempItems[:len(tempItems)-1]

			validUsers, errUsers := db.Query("SELECT * FROM users.users WHERE user_id IN (" + tempItems + ")")
			if errUsers != nil {
				errorManager(w, "No se encontró algún usuario")
				panic(err)
			}
			defer validUsers.Close()
			for validUsers.Next() {
				errUsers = validUsers.Scan(&user.Id, &user.Email)
				if errUsers != nil {
					errorManager(w, "No se encontró algún usuario")
					panic(err)
				}
				users = append(users, user)
			}

			if len(users) != len(scheduled_items.Uuids) {
				errorManager(w, "No se encontró algún usuario")
			} else {

				tempItems = strings.Replace(tempItems, "'", "", -1)
				tempItems = strings.Replace(tempItems, ",", ", ", -1)

				//INSERT
				var insertId string
				insertAlarm := db.QueryRow("INSERT INTO scheduled_items.scheduled_items (id,description,users) VALUES (gen_random_uuid(),'" + scheduled_items.Description + "','{" + tempItems + "}') RETURNING id")
				errUsers := insertAlarm.Scan(&insertId)
				if errUsers != nil {
					errorManager(w, "No se registró")
					panic(errUsers)
				} else {

					scheduled_items.AlarmId = insertId

					go StartTimer(insertId, scheduled_items.AlarmAt)
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(scheduled_items)
				}
			}
		}
	} else {
		errorManager(w, "Alarm should be after now")
	}
}

func errorManager(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}
