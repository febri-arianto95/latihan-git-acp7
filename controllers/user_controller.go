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

	var userDB models.User
	userDB.Name = userInput.Name
	userDB.Email = userInput.Email
	userDB.Password = userInput.Password
	check_email := configs.DB.Where("email = ?", userInput.Email).Find(&userDB).Error
	if check_email != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: "email sudah ada ->" + check_email.Error(),
			Status:  "error",
		})
	} else {
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

	// token, err := middleware.GenerateToken(int(userDB.ID), userDB.Name)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, models.UserResponse{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: err.Error(),
	// 		Status:  "error jwt",
	// 	})
	// }

	// var userTokenResponse = models.UsersTokenResponse{
	// 	ID:    userDB.ID,
	// 	Name:  userDB.Name,
	// 	Email: userDB.Email,
	// 	Token: token,
	// }

	// return c.JSON(http.StatusOK, models.UserTokenResponseSingle{
	// 	Code:    http.StatusOK,
	// 	Message: "Success register",
	// 	Status:  "success",
	// 	Data:    userTokenResponse,
	// })
}

func LoginUsersController(c echo.Context) error {
	var userInput models.UserRequest
	c.Bind(&userInput)

	var userDB models.User

	err := configs.DB.Where("email = ? AND password = ?", userInput.Email, userInput.Password).Find(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}

	token, err := middleware.GenerateToken(int(userDB.ID), userDB.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserTokenResponseSingle{
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
	// categoryId := c.QueryParam("categoryId")
	// page := c.QueryParam("page")
	// userId, _ := strconv.Atoi(c.Param("userId"))
	// fmt.Println("ini token=", c)
	var users []models.User
	err := configs.DB.Find(&users).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
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
