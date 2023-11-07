package migrations

import (
	"saas/infra/database"
	"saas/models"
)

func Migrate() {
	var migrationModels = []interface{}{&models.User{}}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
