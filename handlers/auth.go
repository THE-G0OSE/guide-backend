package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/THE-G0OSE/guide-backend/database"
	"github.com/THE-G0OSE/guide-backend/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var in models.AuthRequest
	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := models.RequestToUser(in)

	if err := database.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusConflict, "user already exists")
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": user.ID})
}

func Login(c echo.Context) error {
	var in models.AuthRequest
	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user models.User

	if err := database.DB.Where("username = ? AND password = ?", in.Username, in.Password).First(&user).Error; err != nil {
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return c.JSON(http.StatusOK, claims)
}
