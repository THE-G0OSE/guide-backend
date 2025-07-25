package main

import (
	"github.com/THE-G0OSE/guide-backend/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/ping", handlers.Ping)

}
