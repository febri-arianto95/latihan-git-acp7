package models

import (
	"time"

	"gorm.io/gorm"
)

//DB transactions
type Transaction struct {
	ID        uint                `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt time.Time           `json:"createdAt", form:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt", form:"updatedAt"`
	DeletedAt gorm.DeletedAt      `json:"deletedAt", form:"deletedAt", gorm:"index"`
	IDUser    uint                `json:"id_user", form:"id_user"`
	Total     uint                `json:"total", form:"total"`
	Status    string              `json:"status", form:"status", gorm:"type:string('checkout','paid','send','delivered')"`
	User      User                `gorm:"foreignKey:IDUser"`
	Detail    []DetailTransaction `gorm:"foreignKey:IDTransaction"`
}

//DB details_transaction
type DetailTransaction struct {
	ID            uint           `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt     time.Time      `json:"createdAt", form:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt", form:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt", form:"deletedAt", gorm:"index"`
	IDTransaction uint           `json:"id_transaction,form;"id_transaction"`
	IDProduct     uint           `json:"id_product", form:"id_product"`
	Quantity      uint           `json:"quantity", form:"quantity"`
	Transactions  Transaction    `gorm:"foreignKey:IDTransaction"`
	Product       Product        `gorm:"foreignKey:IDProduct"`
}
type CheckoutRequest struct {
	IDCart uint `json:"id_cart", form:"id_cart"`
}
type TransactionResponseAny struct {
	Code    int         `json:"code", form:"code"`
	Message string      `json:"message", form:"message"`
	Status  string      `json:"status", form:"status"`
	Data    Transaction `json:"data", form:"data"`
}
type TransactionResponseMany struct {
	Code    int           `json:"code", form:"code"`
	Message string        `json:"message", form:"message"`
	Status  string        `json:"status", form:"status"`
	Data    []Transaction `json:"data", form:"data"`
}
