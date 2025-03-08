package mysql

import (
	"mathgpt/app/models"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	)
}
