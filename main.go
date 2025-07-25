package main

import (
	"os"

	"github.com/THE-G0OSE/guide-backend/database"
	"github.com/THE-G0OSE/guide-backend/handlers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.Connect()
	database.DB.Debug().AutoMigrate()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", handlers.Ping)

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	r := e.Group("/api")
	config := echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	r.Use(echojwt.WithConfig(config))
	r.GET("/me", handlers.Me)

	e.Logger.Fatal(e.Start(":3000"))

}
