package handlers

import (
	"log"
	"net/http"

	"github.com/THE-G0OSE/guide-backend/helpers"
	"github.com/THE-G0OSE/guide-backend/models"
	"github.com/THE-G0OSE/guide-backend/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CourseHandler struct{ Repo *repository.CourseRepo }

func (h *CourseHandler) CreateCourse(c echo.Context) error {

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

	if err := h.Repo.Create(&course); err != nil {
		log.Print("failed to create course")
		return echo.NewHTTPError(http.StatusConflict, "failed to create course")
	}

	return c.JSON(http.StatusCreated, echo.Map{"course": course})
}

func (h CourseHandler) GetCourse(c echo.Context) error {

	id, err := helpers.UintParse(c.Param("id"))

	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse id")
	}

	course, err := h.Repo.Find(id)
	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusBadRequest, "course not found")
	}

	return c.JSON(http.StatusOK, echo.Map{"course": course})

}

func (h CourseHandler) GetMyCourses(c echo.Context) error {

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

	courses, err := h.Repo.FindMy(claims.UserID)
	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, echo.Map{"courses": courses})

}

func (h CourseHandler) DeleteCourse(c echo.Context) error {

	user, ok := c.Get("user").(*jwt.Token)

	if !ok {
		log.Print("failed to get jwt")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	claims, ok := user.Claims.(*models.JwtCustomClaims)

	if !ok {
		log.Print("failed to parse claims")
	}

	id, err := helpers.UintParse(c.Param("id"))

	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse id")
	}

	course, err := h.Repo.Find(id)
	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusNotFound, "course not found")
	}

	if course.CreatorID != claims.UserID {
		return echo.NewHTTPError(http.StatusUnauthorized, "you have not permission to this action")
	}

	if err := h.Repo.Delete(course); err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete course")
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "course was succesfully deleted"})

}

func (h CourseHandler) PatchCourse(c echo.Context) error {
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

	id, err := helpers.UintParse(c.Param("id"))

	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse id from params")
	}

	course, err := h.Repo.Find(id)

	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if claims.UserID != course.CreatorID {
		log.Print("user have no permission to this action")
		return echo.NewHTTPError(http.StatusUnauthorized, "you have no permission to this action")
	}

	var patch models.CoursePathcRequest

	if err := c.Bind(&patch); err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	if patch.Name != nil {
		course.Name = patch.Name
	}
	if patch.Levels != nil {
		course.Levels = patch.Levels
	}

	if err := h.Repo.Patch(course); err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, course)

}
