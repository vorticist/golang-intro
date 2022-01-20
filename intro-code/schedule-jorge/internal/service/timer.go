package service

import (
	"fmt"
	"strings"
	"time"
)

func StartTimer(alarmId string, alarmTime string) {
	ch := make(chan time.Duration)
	defer close(ch)
	go timerFunc(alarmId, ch)

	alarmAtTime, _ := time.Parse("2006-01-02 15:04:05", alarmTime)
	strNow := time.Now().Format("2006-01-02 15:04:05")
	timeNow, _ := time.Parse("2006-01-02 15:04:05", strNow)

	if alarmAtTime.Sub(timeNow).Seconds() > 0 {
		ch <- alarmAtTime.Sub(timeNow)
		fmt.Println(alarmAtTime.Sub(timeNow))
	} else {
		ch <- 1 * time.Second
		fmt.Println(1 * time.Second)
	}
}

func timerFunc(alarmId string, ch chan time.Duration) {
	timer := time.NewTimer(<-ch)
	defer timer.Stop()

	//timer finished
	<-timer.C

	db := getConnection()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM scheduled_items.scheduled_items WHERE id = '" + alarmId + "'")

	var idAlarm string
	var description string
	var users string
	err := row.Scan(&idAlarm, &description, &users)
	if err != nil {
		panic(err)
	} else {
		var user Users
		usersArr := strings.Replace(users, "{", "('", -1)
		usersArr = strings.Replace(usersArr, "}", "')", -1)
		usersArr = strings.Replace(usersArr, ",", "','", -1)

		validUsers, errUsers := db.Query("SELECT * FROM users.users WHERE user_id IN " + usersArr + ";")
		if errUsers != nil {
			panic(errUsers)
		}
		defer validUsers.Close()
		var emailsArr string
		for validUsers.Next() {
			errUsers = validUsers.Scan(&user.Id, &user.Email)
			if errUsers != nil {
				panic(errUsers)
			}
			emailsArr = emailsArr + user.Email + ","
		}

		emailsArr = emailsArr[:len(emailsArr)-1]
		var finishedAlarmId string

		insertAlarmFinished := db.QueryRow("INSERT INTO output.output (id,description,emails) VALUES (gen_random_uuid(),'" + description + "','{" + emailsArr + "}') RETURNING id")
		errFinishedAlarm := insertAlarmFinished.Scan(&finishedAlarmId)
		if errFinishedAlarm != nil {
			panic(errUsers)
		} else {
			fmt.Println("alarm finished: " + finishedAlarmId)
		}
	}
}
