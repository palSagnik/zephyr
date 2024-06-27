package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/palSagnik/zephyr/config"
)

var (
	DB *sql.DB
	err error
)
func ConnectDB() error {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASS, config.DB_NAME)

	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(30)
	DB.SetMaxIdleConns(30)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = DB.PingContext(ctx)
	return err
}
