package main

import (
	"reminders/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo) {
	// For User
	// Route / handler function
	e.GET("/users", handlers.GetUsers)
	e.POST("/users", handlers.NewUser)
	e.PUT("users/:userid:email", handlers.UpdateUser)
	e.DELETE("users/:id", handlers.DeleteUser)
	// For Schedule
	e.GET("/schedules", handlers.GetSchedules)
	e.POST("/schedules", handlers.NewSchedule)
	e.PUT("schedules/:id", handlers.UpdateSchedule)
	e.DELETE("schedules/:id", handlers.DeleteSchedule)
	//For Output
	e.GET("/outputs", handlers.GetSchedules)
	e.POST("/outputs", handlers.NewSchedule)
	//Question: How receive more params?
	e.PUT("outputs/:id/:description/:emails", handlers.UpdateOutput)
	e.DELETE("outputs/:id", handlers.DeleteOutput)
}
