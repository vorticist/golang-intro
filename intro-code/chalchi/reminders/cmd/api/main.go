package main

import (
	"github.com/labstack/echo/v4"
	"reminders/internal/config"
)

func main() {
	// create a new echo instance
	e := echo.New()
	//create Routes
	config.InitRouting(e)	
	e.Logger.Fatal(e.Start(":3000"))
}
