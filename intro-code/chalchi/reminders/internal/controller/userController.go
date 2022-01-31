package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"reminders/internal/models"
	"reminders/internal/repository"
)

func NewUser(c echo.Context) error {
	user := models.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	defer c.Request().Body.Close()
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is your user %v", user)
	//TODO: call the create new user into data base
	g := repository.NewUserRepository()
	userid := g.NewUser(models.User{
		IdUser: repository.GenerateUUID(),
		Email:  user.Email,
	})
	log.Printf("The new user id is %v", userid)
	return c.String(http.StatusOK, "user created successfully")
}

func UpdateUser(c echo.Context) error {
	user := models.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is your user %#v to update", user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	// userId := c.QueryParam("userid")
	dataType := c.Param("data")

	if dataType == "json" {
		// user := vo.User{
		// 	IdUser: userId,
		// 	Email:  "",
		// }
		//TODO: Make the delete user into data base.
		return c.JSON(http.StatusOK, "")
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data",
		})
	}
}

func GetUsers(c echo.Context) error {
	g := repository.NewUserRepository()
	users, _ := g.ListUsers()
	usersToJson, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("users to json %#v\n", string(usersToJson))
	return c.JSON(http.StatusOK, string(usersToJson))
}
