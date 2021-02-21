package database

import (
	"fiber-gorm/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetDBConn initialize db session based on dialector
// Read more at: https://gorm.io/docs/connecting_to_the_database.html
func GetDBConn() (*gorm.DB, error) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", config.Env("DB_HOST"), config.Env("DB_USER"), config.Env("DB_PASS"), config.Env("DB_NAME"), config.Env("DB_PORT"), config.Env("DB_TZ"))
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
