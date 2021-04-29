package models

import (
	"time"

	"gorm.io/gorm"
)

//DB Category
type Category struct {
	ID        uint           `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt", form:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt", form:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt", form:"deletedAt", gorm:"index"`
	Name      string         `json:"name", form:"name"`
	Product   []Product      `gorm:"foreignKey:IDCategory"`
}
type CategoryRequest struct {
	Name string `json:"name", form:"name"`
}
type ProductToCategoryResponse struct {
	ID   uint   `json:"id", form:"id", gorm:"primarykey"`
	Name string `json:"name", form:"name"`
}
type CategoryToProductResponse struct {
	ID      uint   `json:"id", form:"id", gorm:"primarykey"`
	Name    string `json:"name", form:"name"`
	Product []ProductResponse
}
type CategoryResponseAny struct {
	Code    int      `json:"code", form:"code"`
	Message string   `json:"message", form:"message"`
	Status  string   `json:"status", form:"status"`
	Data    Category `json:"data", form:"data"`
}
type CategoryResponseMany struct {
	Code    int        `json:"code", form:"code"`
	Message string     `json:"message", form:"message"`
	Status  string     `json:"status", form:"status"`
	Data    []Category `json:"data", form:"data"`
}
