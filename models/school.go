package models

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	ID        int            `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time      `json:"createdAt", form:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt", form:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt", form:"deletedAt", gorm:"index"`
	Name      string
}
