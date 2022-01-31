package controller

import (
	"encoding/json"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"	
	"net/http"
	"reminders/internal/models"
	"reminders/internal/repository"
	"reminders/internal/service"
)

func NewSchedule(c echo.Context) error {
	schedule := models.Schedule{}
	response := models.ResponseSchedule{}
	err := json.NewDecoder(c.Request().Body).Decode(&schedule)
	defer c.Request().Body.Close()
	if err != nil {
		response.Success = false
		response.Message = "Incorrect information received, remember that the correct date format is: YYYY-MM-DDTHH:MM:SSZ, it's important put the indicators T and Z"
		jsonData, _ := json.Marshal(response)
		return c.String(http.StatusBadRequest, string(jsonData))
	}
	ur := repository.NewUserRepository()
	res := checkUsers(schedule.Users)
	//res := ur.ExistUsers(schedule.Users)
	if res {
		emails := ur.GetEmails(schedule.Users)			 
		if len(emails) > 0 {
			s := repository.NewScheduleRepository()
			scheduleId := s.NewSchedule(models.Schedule{
				Id:          repository.GenerateUUID(),
				Description: schedule.Description,
				Users:       schedule.Users,
			})
			if len(scheduleId) > 0 {
				response.Success = true
				response.Message = "Schedule created successfully!"
				jsonData, _ := json.Marshal(response)
				//Call the services for send the reminders
				service.SendReminder(*schedule.Date, models.Output{Description: schedule.Description, Emails: emails})
				return c.String(http.StatusOK, string(jsonData))
			} else {
				response.Success = false
				response.Message = "Schedule not created"
				jsonData, _ := json.Marshal(response)
				return c.String(http.StatusBadRequest, string(jsonData))
			}
		}
		response.Success = false
		response.Message = "User's emails not found"
		jsonData, _ := json.Marshal(response)
		return c.String(http.StatusBadRequest, string(jsonData))
	} else {
		response.Success = false
		response.Message = "One or more users do not exist"
		jsonData, _ := json.Marshal(response)
		return c.JSON(http.StatusBadRequest, string(jsonData))
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
		schedule := models.Schedule{}
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
func contains(s []models.User, str uuid.UUID) bool {
	var r bool = false
	for _, v := range s {
		if v.IdUser == str {
			r = true
			break
		}
	}
	return r
}
