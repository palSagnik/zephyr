package database

import (
	"log"

	"github.com/palSagnik/zephyr/models"
)

var (
	users = new(models.User)
	configs = new(models.CustomConfigurations)
	runningInstances = new(models.RunningInstance)
	verification = new(models.Verification)
)
func MigrateUp() error {

	// creating tables
	err := DB.AutoMigrate(users, configs, verification)
	if err != nil {
		return err
	}
	return nil
}

func MigrateDown(choice string) error {
	
	switch choice {
	case "users":
		err := DB.Migrator().DropTable(users)
		if err != nil {
			return err
		}
		log.Println("Dropped users table")
	case "configs":
		err := DB.Migrator().DropTable(configs)
		if err != nil {
			return err
		}
		log.Println("Dropped configs table")
	case "running":
		err := DB.Migrator().DropTable(runningInstances)
		if err != nil {
			return err
		}
		log.Println("Dropped running instances table")
	case "toverify":
		err := DB.Migrator().DropTable(verification)
		if err != nil {
			return err
		}
		log.Println("Dropped verification table table")
	}

	return nil
}