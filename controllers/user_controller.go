package controllers

import (
	"net/http"
	"project/configs"
	"project/middleware"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

func CreateUsersController(c echo.Context) error {
	var userInput models.UserRequest
	c.Bind(&userInput)
	if userInput.Email == "" || userInput.Password == "" || userInput.Name == "" {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: "All input is required",
			Status:  "error",
		})
	}

	var userDB models.User
	userDB.Name = userInput.Name
	userDB.Email = userInput.Email
	userDB.Password = userInput.Password
	check_email := configs.DB.Where("email = ?", userInput.Email).Find(&userDB).RowsAffected
	if check_email != 0 {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: "Email is already exist",
			Status:  "error",
		})
	}
	err := configs.DB.Save(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}
	return LoginUsersController(c)
}

func LoginUsersController(c echo.Context) error {
	var userInput models.UserRequest
	c.Bind(&userInput)
	var userDB models.User
	check_user := configs.DB.Where("email = ? AND password=?", userInput.Email, userInput.Password).Find(&userDB).RowsAffected
	if check_user == 0 {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: "Invalid email or password",
			Status:  "error",
		})
	}
	err := configs.DB.Where("Email = ? AND Password = ?", userInput.Email, userInput.Password).Find(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}
	token, err := middleware.GenerateToken(int(userDB.ID), userDB.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error jwt",
		})
	}

	var userTokenResponse = models.TokenResponse{
		ID:    userDB.ID,
		Name:  userDB.Name,
		Email: userDB.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, models.UserTokenResponseSingle{
		Code:    http.StatusOK,
		Message: "Success Login",
		Status:  "success",
		Data:    userTokenResponse,
	})
}

func GetUsersController(c echo.Context) error {
	var users []models.User
	err := configs.DB.Find(&users).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}

	// userId dari JWT
	userId := middleware.ExtractUserIdFromJWT(c)

	return c.JSON(http.StatusOK, models.UserResponse{
		Code:    http.StatusOK,
		Message: "Success get data user id= " + strconv.Itoa(userId),
		Status:  "success",
		Data:    users,
	})
}
