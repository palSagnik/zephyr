package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/palSagnik/zephyr/config"
)


func ConnectDB() error {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASS, config.DB_NAME)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	return nil
}