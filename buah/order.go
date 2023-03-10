package buah

import (
	"ProjectBuahIn/models"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User      models.User `gorm:"foreignkey:UserID"`
	Product   Buah        `gorm:"foreignkey:ProductID"`
	UserID    uint
	ProductID uint
	Quantity  int `json:"quantity"`
}
