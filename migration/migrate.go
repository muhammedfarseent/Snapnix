package migration

import (
	"log"

	"SNAPNIX/utils/models"

	"gorm.io/gorm"
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
