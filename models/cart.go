package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt", form:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt", form:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt", form:"deletedAt", gorm:"index"`
	IDUser    uint           `json:"id_user", form:"id_user"`
	IDProduct uint           `json:"id_product", form:"id_product"`
	Quantity  uint           `json:"quantity", form:"quantity"`
	Users     User           `gorm:"foreignKey:IDUser"`
	Product   Product        `gorm:"foreignKey:IDProduct"`
}
type CartRequest struct {
	IDProduct uint `json:"id_product", form:"id_product"`
	Quantity  uint `json:"quantity", form:"quantity"`
}
type CartResponseAny struct {
	Code    int    `json:"code", form:"code"`
	Message string `json:"message", form:"message"`
	Status  string `json:"status", form:"status"`
	Data    Cart   `json:"data", form:"data"`
}
type CartResponseMany struct {
	Code    int    `json:"code", form:"code"`
	Message string `json:"message", form:"message"`
	Status  string `json:"status", form:"status"`
	Data    []Cart `json:"data", form:"data"`
}
