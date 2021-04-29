package controllers

import (
	"net/http"
	"project/configs"
	"project/middleware"
	"project/models"

	"github.com/labstack/echo"
	"gorm.io/gorm/clause"
)

func CreateCartController(c echo.Context) error {
	userId := middleware.ExtractUserIdFromJWT(c)
	var cartInput models.CartRequest
	c.Bind(&cartInput)

	//data user
	var userDB models.User
	err_userDB := configs.DB.Find(&userDB, userId).Error
	if err_userDB != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_userDB.Error(),
			Status:  "error",
		})
	}

	//data product

	// err_categories := configs.DB.Find(&categories, "id_category=?", cartInput.IDProduct)

	var productDB models.Product

	err_productDB := configs.DB.Find(&productDB, cartInput.IDProduct).Error
	if err_productDB != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_productDB.Error(),
			Status:  "error",
		})
	}

	// var categoryDB models.Category
	// err_categoryDB := configs.DB.Find(&categoryDB, productDB.Category.ID)

	// if err_categoryDB != nil {
	// 	return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: err_productDB.Error(),
	// 		Status:  "error",
	// 	})
	// }

	//data cart
	var cartDB models.Cart
	cartDB.IDUser = uint(userId)
	cartDB.IDProduct = cartInput.IDProduct
	cartDB.Quantity = cartInput.Quantity
	cartDB.Users = userDB
	cartDB.Product = productDB

	cartDB.Product.Category.ID = cartDB.Product.IDCategory
	cartDB.Product.Category.Name = cartDB.Product.Name
	cartDB.Product.Category.CreatedAt = cartDB.Product.CreatedAt
	cartDB.Product.Category.UpdatedAt = cartDB.Product.UpdatedAt
	cartDB.Product.Category.DeletedAt = cartDB.Product.DeletedAt
	// fmt.Println(cartDB.Product.Category.ID)

	err := configs.DB.Save(&cartDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  "error",
		})
	}
	return c.JSON(http.StatusOK, models.CartResponseAny{
		Code:    http.StatusOK,
		Message: "Success add cart",
		Status:  "success",
		Data:    cartDB,
	})
}
func GetCartController(c echo.Context) error {
	userId := middleware.ExtractUserIdFromJWT(c)
	var cartDB []models.Cart

	err_cart := configs.DB.Where("id_user = ?", userId).Preload(clause.Associations).Find(&cartDB).Error

	if err_cart != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_cart.Error(),
			Status:  "error",
		})
	}
	return c.JSON(http.StatusOK, models.CartResponseMany{
		Code:    http.StatusOK,
		Message: "Success get data all cart",
		Status:  "success",
		Data:    cartDB,
	})
}
func DeleteCartController(c echo.Context) error {
	userId := middleware.ExtractUserIdFromJWT(c)
	cartId := c.Param("id")
	var cartDB []models.Cart
	err_cart := configs.DB.Where("id = ? AND id_user = ?", cartId, userId).Delete(&cartDB).Error
	if err_cart != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_cart.Error(),
			Status:  "error",
		})
	}
	return c.JSON(http.StatusOK, models.ResponseNotif{
		Code:    http.StatusOK,
		Message: "Deleted success",
		Status:  "success",
	})
}
