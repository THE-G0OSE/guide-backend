package main

import (
	"os"

	"github.com/THE-G0OSE/guide-backend/database"
	"github.com/THE-G0OSE/guide-backend/handlers"
	"github.com/THE-G0OSE/guide-backend/models"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.Connect()
	database.DB.Debug().AutoMigrate(&models.User{})

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", handlers.Ping)

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	r := e.Group("/api")
	config := echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
	}

	r.Use(echojwt.WithConfig(config))
	r.GET("/me", handlers.Me)

	e.Logger.Fatal(e.Start(":3000"))

}
