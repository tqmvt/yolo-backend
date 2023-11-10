package migrations

import (
	"saas/infra/database"
	"saas/models"
)

func Migrate() {
	var migrationModels = []interface{}{&models.User{}, &models.Round{}, &models.Player{}}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
