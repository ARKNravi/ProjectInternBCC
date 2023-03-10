package buah

import (
	"time"

	"gorm.io/gorm"
)

type BuahRequest struct {
	gorm.Model
	Nama        string `json:"Nama" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint   `json:"price" binding:"required,number"`
	Discount    uint   `json:"discount" binding:"required,number"`
	Quantity    uint   `json:"quantity" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
