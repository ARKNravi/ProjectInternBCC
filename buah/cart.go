package buah

import (
	"ProjectBuahIn/models"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	User       models.User `gorm:"foreignkey:UserID"`
	Product    Buah        `gorm:"foreignkey:ProductID"`
	Quantity   uint        `json:"quantity"`
	TotalPrice uint        `json:"totalprice"`
}
