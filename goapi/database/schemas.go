package database

import (
	"context"
	"log"
	"time"
)

func CreateTables() error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err = DB.QueryContext(ctx, `
	CREATE TABLE IF NOT EXISTS users(
		userid bigserial PRIMARY KEY,
		email text NOT NULL UNIQUE,
		username text NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL,
	)`)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = DB.QueryContext(ctx, `
	CREATE TABLE IF NOT EXISTS running(
		runid bigserial PRIMARY KEY,
		userid bigint NOT NULL REFERENCES users(userid),
		`)
	if err != nil {
		log.Println(err)
		return err
	}
	
	return nil
}