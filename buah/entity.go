package buah

import (
	"time"

	"gorm.io/gorm"
)

type Buah struct {
	gorm.Model
	Nama        string
	Description string
	Price       uint
	Discount    uint
	Quantity    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
