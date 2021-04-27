package controllers

import (
	"net/http"
	"project/configs"
	"project/models"

	"github.com/labstack/echo"
	"gorm.io/gorm/clause"
)

func CreateProductController(c echo.Context) error {
	var productInput models.ProductRequest
	c.Bind(&productInput)

	var categories models.Category
	err_categories := configs.DB.Find(&categories, productInput.IDCategory).Error

	if err_categories != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err_categories.Error(),
			Status:  "error",
		})
	}
	var productDB models.Product
	productDB.IDCategory = productInput.IDCategory
	productDB.Name = productInput.Name
	productDB.Description = productInput.Description
	productDB.Stock = productInput.Stock
	productDB.Price = productInput.Price
	productDB.Category = categories
	err := configs.DB.Save(&productDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ProductResponseAny{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}

	return c.JSON(http.StatusOK, models.ProductResponseAny{
		Code:    http.StatusOK,
		Message: "Success add product",
		Status:  "success",
		Data:    productDB,
	})
}
func GetProductController(c echo.Context) error {
	// categoryId := c.QueryParam("categoryId")
	// var temp uint
	// temp=strconv.ParseUint(categoryId, 10, 0)
	// var categoryDB models.Category
	// err_categories := configs.DB.Find(&categoryDB, ).Error

	// if err_categories != nil {

	// }
	var productDB []models.Product
	// err := configs.DB.Joins("Category").Find(&productDB).Error
	err := configs.DB.Preload(clause.Associations).Find(&productDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ProductResponseMany{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}
	return c.JSON(http.StatusOK, models.ProductResponseMany{
		Code:    http.StatusOK,
		Message: "Success get data all product",
		Status:  "success",
		Data:    productDB,
	})
}

func GetProByCatController(c echo.Context) error {
	var categoryDB []models.Category
	err := configs.DB.Preload(clause.Associations).Find(&categoryDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.CategoryResponseMany{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}
	return c.JSON(http.StatusOK, models.CategoryResponseMany{
		Code:    http.StatusOK,
		Message: "Success get data all product",
		Status:  "success",
		Data:    categoryDB,
	})
}
