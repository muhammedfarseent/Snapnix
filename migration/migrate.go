package migration

import (
	"log"

	"gorm.io/gorm"

	"ECOMERS/utils/models"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Fatal("failed to migrate")
	}
	log.Println("migration sucessfull")
}
