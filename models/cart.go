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
