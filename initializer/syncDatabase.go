package initializer

import "ProjectBuahIn/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
