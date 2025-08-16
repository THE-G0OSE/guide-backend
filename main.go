package main

import (
	"os"

	"github.com/THE-G0OSE/guide-backend/database"
	"github.com/THE-G0OSE/guide-backend/handlers"
	"github.com/THE-G0OSE/guide-backend/models"
	"github.com/THE-G0OSE/guide-backend/repository"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.Connect()
	database.DB.Debug().AutoMigrate(&models.User{}, &models.Course{}, &models.Level{}, &models.Block{})

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", handlers.Ping)

	userRepo := &repository.UserRepo{DB: database.DB}
	authHandler := &handlers.AuthHandler{Repo: userRepo}

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	r := e.Group("/api")
	config := echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
	}

	r.Use(echojwt.WithConfig(config))
	r.GET("/me", handlers.Me)

	course := r.Group("/courses")
	courseRepo := &repository.CourseRepo{DB: database.DB}
	courseHandler := &handlers.CourseHandler{Repo: courseRepo}

	course.POST("/create", courseHandler.CreateCourse)
	course.GET("/getone/:id", courseHandler.GetCourse)
	course.GET("/getmy", courseHandler.GetMyCourses)
	course.DELETE("/delete/:id", courseHandler.DeleteCourse)
	course.PATCH("/patch/:id", courseHandler.PatchCourse)

	e.Logger.Fatal(e.Start(":3000"))

}
