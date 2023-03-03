package buah

import "time"

type Buah struct {
	ID          int
	Jenis       string
	Description string
	Price       int
	Discount    int
	Matang      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
