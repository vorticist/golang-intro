package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"reminders/cmd/api/vo"
	"reminders/cmd/cron"
	"reminders/internal/repository"

	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func NewSchedule(c echo.Context) error {
	schedule := vo.Schedule{}
	err := json.NewDecoder(c.Request().Body).Decode(&schedule)
	defer c.Request().Body.Close()
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	ur := repository.NewUserRepository()
	res := checkUsers(schedule.Users)
	//res := ur.ExistUsers(schedule.Users)
	if res {
		emails := ur.GetEmails(schedule.Users)
		cron.WriteToOutput(emails)
		s := repository.NewScheduleRepository()
		scheduleId := s.NewSchedule(repository.Schedule{
			Id:          repository.GenerateUUID(),
			Description: schedule.Description,
			Users:       schedule.Users,
		})
		log.Printf("The new schedule id is %v", scheduleId)
		//TODO: validate if it was created correctly
		//Write the channel
		return c.String(http.StatusOK, "Schedule created successfully!")
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "One or more users do not exist",
		})
	}
}

func UpdateSchedule(c echo.Context) error {
	// id := c.QueryParam("id")
	// description := c.QueryParam("description")
	// users := c.QueryParam("users")
	dataType := c.Param("data")

	if dataType == "json" {
		// schedule := vo.Schedule{
		// 	Id:          id,
		// 	Description: description,
		// 	Users:       users,
		// }
		//TODO: Make the update schedule into data base.
		return c.JSON(http.StatusOK, "")
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data for the update user",
		})
	}
}

func DeleteSchedule(c echo.Context) error {
	//id := c.QueryParam("id")
	dataType := c.Param("data")

	if dataType == "json" {
		// schedule := vo.Schedule{
		// 	Id:          id,
		// 	Description: "",
		// 	Users:       "",
		// }
		//TODO: Make the delete schedule into data base.
		return c.JSON(http.StatusOK, "")
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data",
		})
	}
}

func GetSchedules(c echo.Context) error {
	dataType := c.Param("data")
	if dataType == "json" {
		schedule := vo.Schedule{}
		return c.JSON(http.StatusOK, schedule)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Any Error",
		})
	}
}

//TODO: move this logic in other layer
func checkUsers(users []uuid.UUID) bool {
	ur := repository.NewUserRepository()
	usersDb, _ := ur.ListUsers()
	var result bool = false
	for _, u := range users {
		result = contains(usersDb, u)
		if !result {
			break
		}
	}
	return result
}

//TODO: Investigate if exit other better form
func contains(s []repository.User, str uuid.UUID) bool {
	var r bool = false
	for _, v := range s {
		if v.IdUser == str {
			r = true
			break
		}
	}
	return r
}
