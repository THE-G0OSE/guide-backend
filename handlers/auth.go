package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/THE-G0OSE/guide-backend/models"
	"github.com/THE-G0OSE/guide-backend/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Repo *repository.UserRepo
}

func (h AuthHandler) Register(c echo.Context) error {
	var in models.AuthRequest
	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := models.RequestToUser(in)

	if err := h.Repo.Create(&user); err != nil {
		return echo.NewHTTPError(http.StatusConflict, "user already exists")
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": user.ID})
}

func (h AuthHandler) Login(c echo.Context) error {

	var in models.AuthRequest

	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.Repo.Login(in.Password, in.Username)
	if err != nil {
		return echo.ErrUnauthorized
	}

	claims := &models.JwtCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"token": t})
}

func Me(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "no token")
	}
	claims, ok := user.Claims.(*models.JwtCustomClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid claims")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"uid":      claims.UserID,
		"username": claims.Username,
	})
}
