package models

import (
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Activity{},
		&DataSource{},
		&Commit{},
		&Repository{},
	)
}