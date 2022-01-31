package config

import (
	"github.com/labstack/echo/v4"
	"reminders/internal/controller"
)

func InitRouting(e *echo.Echo) {
	// For User
	// Route / handler function
	e.GET("/users", controller.GetUsers)
	e.POST("/users", controller.NewUser)
	//e.PUT("users/:userid:email", controller.UpdateUser)
	//e.DELETE("users/:id", controller.DeleteUser)
	// For Schedule
	e.GET("/schedules", controller.GetSchedules)
	e.POST("/schedules", controller.NewSchedule)
	//e.PUT("schedules/:id", controller.UpdateSchedule)
	//e.DELETE("schedules/:id", controller.DeleteSchedule)
	//For Output
	e.GET("/outputs", controller.GetSchedules)
	//e.POST("/outputs", controller.NewSchedule)
	//e.PUT("outputs/:id/:description/:emails", controller.UpdateOutput)
	e.DELETE("outputs/:id", controller.DeleteOutput)
}
