package buah

import "time"

type BuahRequest struct {
	Jenis       string `json:"jenis" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Discount    int    `json:"discount" binding:"required,number"`
	Matang      bool   `json:"Matang" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
