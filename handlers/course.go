package handlers

import (
	"log"
	"net/http"

	"github.com/THE-G0OSE/guide-backend/database"
	"github.com/THE-G0OSE/guide-backend/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateCourse(c echo.Context) error {

	user, ok := c.Get("user").(*jwt.Token)

	if !ok {
		log.Print("failed to get jwt")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	var in models.CourseCreateRequest

	if err := c.Bind(&in); err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	claims, ok := user.Claims.(*models.JwtCustomClaims)

	if !ok {
		log.Print("failed to parse claims")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	course := models.RequestToCourse(in, claims.UserID)

	if err := database.DB.Create(&course).Error; err != nil {
		log.Print("failed to create course")
		return echo.NewHTTPError(http.StatusConflict, "failed to create course")
	}

	return c.JSON(http.StatusCreated, echo.Map{"course": course})
}

func GetCourse(c echo.Context) error {

	id := c.Param("id")

	var course models.Course

	if err := database.DB.Where("ID = ?", id).First(&course).Error; err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusBadRequest, "course not found")
	}

	return c.JSON(http.StatusOK, echo.Map{"course": course})

}

func GetMyCourses(c echo.Context) error {

	user, ok := c.Get("user").(*jwt.Token)

	if !ok {
		log.Print("failed to get jwt")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	claims, ok := user.Claims.(*models.JwtCustomClaims)

	if !ok {
		log.Print("failed to parse claims")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	var courses []models.Course

	if err := database.DB.Where("Creator_ID = ?", claims.UserID).Find(&courses).Error; err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, echo.Map{"courses": courses})

}

func DeleteCourse(c echo.Context) error {

	user, ok := c.Get("user").(*jwt.Token)

	if !ok {
		log.Print("failed to get jwt")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	claims, ok := user.Claims.(*models.JwtCustomClaims)

	if !ok {
		log.Print("failed to parse claims")
	}

	id := c.Param("id")

	var course models.Course

	if err := database.DB.Where("ID = ?", id).First(&course).Error; err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusNotFound, "course not found")
	}

	if course.ID != claims.UserID {
		return echo.NewHTTPError(http.StatusUnauthorized, "you have not permission to this action")
	}

	if err := database.DB.Delete(&course, "ID = ?", course.ID).Error; err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete course")
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "course was succesfully deleted"})

}
