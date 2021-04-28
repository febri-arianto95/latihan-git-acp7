package controllers

import (
	"fmt"
	"net/http"
	"project/configs"
	"project/middleware"
	"project/models"

	"github.com/labstack/echo"
)

func CreateTransactionController(c echo.Context) error {
	userId := middleware.ExtractUserIdFromJWT(c)
	// status := "checkout"
	var input models.CheckoutRequest
	c.Bind(&input)
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

	//data cart
	var cartDB models.Cart
	err_cartDB := configs.DB.Find(&cartDB, input.IDCart).Error
	if err_cartDB != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_cartDB.Error(),
			Status:  "error",
		})
	}

	//data product
	var productDB models.Product
	err_productDB := configs.DB.Find(&productDB, cartDB.IDProduct).Error
	if err_productDB != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_productDB.Error(),
			Status:  "error",
		})
	}
	// total := cartDB.Quantity * productDB.Price

	//max id transaction
	var maxId uint
	row := configs.DB.Table("Transaction").Select("max(ID)").Row()
	row.Scan(&maxId)
	// var trxDB models.Transaction
	// err_trxDB := configs.DB.Select("max(ID)").Find(&trxDB).Error
	// if err_trxDB != nil {
	// 	return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: err_productDB.Error(),
	// 		Status:  "error",
	// 	})
	// }
	//end of max id transaction
	fmt.Println(maxId)
	useId := maxId + 1
	fmt.Println(useId)
	//save data transaction
	// var transactionDB models.Transaction
	// // transactionDB.ID = useId
	// transactionDB.IDUser = uint(userId)
	// transactionDB.Status = status
	// transactionDB.Total = total
	// transactionDB.User = userDB
	// // transactionDB.Detail = detTrxDB

	// if err_trx := configs.DB.Save(&transactionDB).Error; err_trx != nil {
	// 	return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: err_trx.Error(),
	// 		Status:  "error",
	// 	})
	// }
	//save detail transaksi
	var detailTrxDB models.DetailTransaction
	detailTrxDB.IDTransaction = 7
	detailTrxDB.IDProduct = 5
	detailTrxDB.Quantity = 2
	// detailTrxDB.Transactions = transactionDB
	// detailTrxDB.Product = productDB
	if err_detailTrxDB := configs.DB.Save(&detailTrxDB).Error; err_detailTrxDB != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
			Code:    http.StatusInternalServerError,
			Message: err_detailTrxDB.Error(),
			Status:  "error",
		})
	}

	//load multi data detail transaction
	// var detTrxDB []models.DetailTransaction
	// err_detTrxDB := configs.DB.Find(&detTrxDB).Error
	// if err_detTrxDB != nil {
	// 	return c.JSON(http.StatusInternalServerError, models.ResponseNotif{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: err_productDB.Error(),
	// 		Status:  "error",
	// 	})
	// }
	//end multi data detail transaction

	//menambahkan load detail transaksi
	// var loadTrxDB models.Transaction
	// configs.DB.Find(&loadTrxDB, useId)
	// loadTrxDB.Detail = detTrxDB
	// configs.DB.Save(&loadTrxDB)
	//end menambahkan load detail transaksi

	//pengurangan stock
	// configs.DB.Find(&productDB, cartDB.IDProduct)
	// productDB.Stock = productDB.Stock - cartDB.Quantity
	// configs.DB.Save(&productDB)
	//end pengurangan stok

	return c.JSON(http.StatusOK, models.TransactionResponseAny{
		Code:    http.StatusOK,
		Message: "Success add checkout",
		Status:  "success",
		// Data:    transactionDB,
	})
}
