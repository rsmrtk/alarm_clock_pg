package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "C:/Users/Rostyk/OneDrive/Рабочий стол/alarm_clock_pg/static")
	// A handler for rendering an HTML file
	e.GET("/", func(c echo.Context) error {
		return c.File("C:/Users/Rostyk/OneDrive/Рабочий стол/alarm_clock_pg/static/crm.html")
	})

	// Starting the server on port 8888
	e.Logger.Fatal(e.Start(":8888"))
}
