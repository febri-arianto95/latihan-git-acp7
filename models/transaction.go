package models

import (
	"time"

	"gorm.io/gorm"
)

//DB transactions
type Transaction struct {
	ID        uint           `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt", form:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt", form:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt", form:"deletedAt", gorm:"index"`
	IDUser    uint           `json:"id_user", form:"id_user"`
	Status    string         `json:"status", form:"status", gorm:"type:string('checkout','paid','send','delivered')"`
	Users     User           `gorm:"foreignKey:IDUser"`
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
type TransactionsRequest struct {
	IDUser    string `json:"id_user", form:"id_user"`
	IDProduct uint   `json:"id_product", form:"id_product"`
	Quantity  uint   `json:"quantity", form:"quantity"`
}
type TransactionsDBResponse struct {
	ID        uint                `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt time.Time           `json:"createdAt", form:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt", form:"updatedAt"`
	DeletedAt gorm.DeletedAt      `json:"deletedAt", form:"deletedAt", gorm:"index"`
	IDUser    uint                `json:"id_user", form:"id_user"`
	Status    string              `json:"status", form:"status"`
	User      User                `json:"user", form:"user"`
	Product   []DetailTransaction `json:"product", form:"product"`
}
type TransactionsResponseSingle struct {
	Code    int                    `json:"code", form:"code"`
	Message string                 `json:"message", form:"message"`
	Data    TransactionsDBResponse `json:"data", form:"data"`
	Status  string                 `json:"status", form:"status"`
}
type TransactionsResponse struct {
	Code    int                      `json:"code", form:"code"`
	Message string                   `json:"message", form:"message"`
	Data    []TransactionsDBResponse `json:"data", form:"data"`
	Status  string                   `json:"status", form:"status"`
}
