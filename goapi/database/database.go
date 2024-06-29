package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/palSagnik/zephyr/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)
func ConnectDB() error {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASS, config.DB_NAME)

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Printf("Connected To Database: %s\n", config.DB_NAME)
	return nil
}
