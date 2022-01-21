package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"reminders/cmd/api/vo"
	"reminders/internal/repository"

	"github.com/labstack/echo/v4"
)

func NewOutput(c echo.Context) error {
	output := vo.Output{}
	err := json.NewDecoder(c.Request().Body).Decode(&output)
	defer c.Request().Body.Close()
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is your output %#v", output)
	//TODO: Make create the new output into data base.
	or := repository.NewOutputRepository()
	outId := or.NewOutput(
		repository.Output{
			Id:          repository.GenerateUUID(),
			Description: output.Description,
			Emails:      output.Emails,
		})
	log.Printf("The new output id is %v", outId)
	return c.String(http.StatusOK, "output created successfully!")
}

func UpdateOutput(c echo.Context) error {
	// id := c.QueryParam("id")
	// description := c.QueryParam("description")
	// emails := c.QueryParam("emails")
	dataType := c.Param("data")

	if dataType == "json" {
		// output := vo.Output{
		// 	Id:          id,
		// 	Description: description,
		// 	Emails:       emails,
		// }
		//TODO: Make the update output into data base.
		return c.JSON(http.StatusOK, "")
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data for the update user",
		})
	}
}

func DeleteOutput(c echo.Context) error {
	//id := c.QueryParam("id")
	dataType := c.Param("data")

	if dataType == "json" {
		// output := vo.Output{
		// 	Id:          id,
		// 	Description: "",
		// 	Emails:       "",
		// }
		//TODO: Make the delete output into data base.
		return c.JSON(http.StatusOK, "")
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data",
		})
	}
}

func GetOutputs(c echo.Context) error {
	dataType := c.Param("data")
	if dataType == "json" {
		//output := []vo.output; // Call the data base for get all users
		output := vo.Output{}
		return c.JSON(http.StatusOK, output)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Any Error",
		})
	}
}
