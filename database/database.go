package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"SNAPNIX/config"
)

var DB *gorm.DB

func ConnectDB(config config.Config) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("%s,%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBName)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DbPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect with database:%v", err)
	}
	DB = db
	return DB, err

}
