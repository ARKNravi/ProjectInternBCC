package buah

import "time"

type BuahResponse struct {
	ID          int    `json:"id"`
	Jenis       string `json:"jenis"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Matang      bool   `json:"matang"`
	Discount    int    `json:"discount"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
