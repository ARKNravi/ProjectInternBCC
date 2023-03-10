package buah

import (
	"time"

	"gorm.io/gorm"
)

type BuahResponse struct {
	gorm.Model
	Nama        string `json:"nama"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Discount    uint   `json:"discount"`
	Quantity    uint   `json:"quantity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
