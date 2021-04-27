package controllers

import (
	"net/http"
	"project/configs"
	"project/models"

	"github.com/labstack/echo"
)

func CreateProductController(c echo.Context) error {
	var productInput models.ProductRequest
	c.Bind(&productInput)
	var productDB models.Product
	productDB.Name = productInput.Name
	productDB.Description = productInput.Description
	productDB.Stock = productInput.Stock
	productDB.Price = productInput.Price
	productDB.IDCategory = productInput.IDCategory
	err := configs.DB.Save(&productDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ProductResponseAny{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}
	var categories models.Category
	err_categories := configs.DB.Find(&categories).Error

	if err_categories != nil {
		return c.JSON(http.StatusInternalServerError, models.UserResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}
	// var categoryResponse = models.ProductToCategoryResponse{
	// 	ID:   categories.ID,
	// 	Name: categories.Name,
	// }
	var productResponse = models.ProductResponse{
		ID:          productDB.ID,
		IDCategory:  productDB.IDCategory,
		Name:        productDB.Name,
		Description: productDB.Description,
		Stock:       productDB.Stock,
		Price:       productDB.Price,
		Category:    categories,
	}
	return c.JSON(http.StatusOK, models.ProductResponseAny{
		Code:    http.StatusOK,
		Message: "Success add product",
		Status:  "success",
		Data:    productResponse,
	})
}
