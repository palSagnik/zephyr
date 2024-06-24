package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


func ConnectDB() error {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}


}