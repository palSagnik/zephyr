package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/palSagnik/zephyr/config"
	"github.com/palSagnik/zephyr/database"
	"github.com/palSagnik/zephyr/logs"
	"github.com/palSagnik/zephyr/router"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	
	// Logging
	logs.CustomLogger()

	// Connecting to the database
	fmt.Println("Connecting to DB...")
	for {
		if err := database.ConnectDB(); err != nil {
			log.Println(err)
			log.Println("Waiting for a minute")
			time.Sleep(time.Minute * 1)
			continue
		}
		break
	}

	// creating tables
	if err := database.MigrateUp(); err != nil {
		log.Fatal(err)
	}

	// Setting up logger files
	errorLogFile, err := os.OpenFile("./logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer errorLogFile.Close()

	errorWriter := io.MultiWriter(os.Stdout, errorLogFile)
	loggerConfig := logger.Config{Output: errorWriter}


	// Initialize *fiber.App
	app := fiber.New()
	app.Use(logger.New(loggerConfig)) 
	app.Use(recover.New())            				// Prevent process exit due to Fatal()
	router.SetUpRoutes(app)           

	log.Fatal(app.Listen(config.APP_PORT))

}