package migrations

import (
	"github.com/wallace-jr/aguia-branca/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(models.Trip{})
}