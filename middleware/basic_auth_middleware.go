package middleware

import (
	"project/configs"
	"project/models"

	"github.com/labstack/echo"
)

func AuthBasicMiddleware(email, password string, c echo.Context) (bool, error) {
	var user models.User
	err := configs.DB.Where("email = ? AND password = ?", email, password).Find(&user).Error

	if err != nil && user.Email == "" {
		return false, err
	}

	return true, nil
}
