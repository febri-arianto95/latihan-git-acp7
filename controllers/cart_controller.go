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
	var users models.User
	err_users := configs.DB.Find(&users, userId).Error
	if err_users != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_users.Error(),
			Status:  "error",
		})
	}

	//data product
	var products models.Product
	err_products := configs.DB.Find(&products, cartInput.IDProduct).Error
	if err_products != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_users.Error(),
			Status:  "error",
		})
	}

	//data cart
	var cartDB models.Cart
	cartDB.IDUser = uint(userId)
	cartDB.IDProduct = cartInput.IDProduct
	cartDB.Quantity = cartInput.Quantity
	cartDB.Users = users
	cartDB.Product = products

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
