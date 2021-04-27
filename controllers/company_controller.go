package controllers

import (
	"net/http"
	"project/configs"
	"project/models"

	"github.com/labstack/echo"
)

func GetUsersCompanyController(c echo.Context) error {

	var userCompanies []models.UserCompany
	err := configs.DB.Preload("companies").Find(&userCompanies).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}

	return c.JSON(http.StatusOK, models.UserCompanyResponse{
		Code:    http.StatusOK,
		Message: "Success get data user companies",
		Status:  "success",
		Data:    userCompanies,
	})
}

func CreateUserCompaniesController(c echo.Context) error {

	var userCompanyInput models.UserCompanyRequest
	c.Bind(&userCompanyInput)

	var userCompanyDB models.UserCompany
	userCompanyDB.Name = userCompanyInput.Name
	userCompanyDB.CompanyID = userCompanyInput.CompanyID

	err := configs.DB.Save(&userCompanyDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}

	return c.JSON(http.StatusOK, models.UserResponseCompanySingle{
		Code:    http.StatusOK,
		Message: "Success get data user",
		Status:  "success",
		Data:    userCompanyDB,
	})
}
