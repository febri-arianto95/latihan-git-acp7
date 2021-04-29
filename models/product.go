package models

import (
	"time"

	"gorm.io/gorm"
)

//DB Products
type Product struct {
	ID          uint           `json:"id", form:"id", gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt", form:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt", form:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt", form:"deletedAt", gorm:"index"`
	IDCategory  uint           `json:"id_category", form:"id_category"`
	Name        string         `json:"name", form:"name"`
	Description string         `json:"description", form:"description"`
	Stock       uint           `json:"stock", form:"stock"`
	Price       uint           `json:"price", form:"price"`
	Category    Category       `gorm:"foreignKey:IDCategory"`
}
type ProductRequest struct {
	IDCategory  uint   `json:"id_category", form:"id_category"`
	Name        string `json:"name", form:"name"`
	Description string `json:"description", form:"description"`
	Stock       uint   `json:"stock", form:"stock"`
	Price       uint   `json:"price", form:"price"`
	Category    Category
}
type ProductResponse struct {
	ID          uint      `json:"id", form:"id"`
	CreatedAt   time.Time `json:"createdAt", form:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt", form:"updatedAt"`
	IDCategory  uint      `json:"id_category", form:"id_category"`
	Name        string    `json:"name", form:"name"`
	Description string    `json:"description", form:"description"`
	Stock       uint      `json:"stock", form:"stock"`
	Price       uint      `json:"price", form:"price"`
	Category    Category
}
type ProductResponseAny struct {
	Code    int     `json:"code", form:"code"`
	Message string  `json:"message", form:"message"`
	Status  string  `json:"status", form:"status"`
	Data    Product `json:"data", form:"data"`
}
type ProductResponseMany struct {
	Code    int       `json:"code", form:"code"`
	Message string    `json:"message", form:"message"`
	Status  string    `json:"status", form:"status"`
	Data    []Product `json:"data", form:"data"`
}
