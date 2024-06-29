package database

import (
	"github.com/palSagnik/zephyr/models"
)

func MigrateUp() error {
	var err error
	
	// struct pointers
	users := new(models.User)
	configs := new(models.CustomConfigurations)
	runningInstances := new(models.RunningInstance)
	verification := new(models.Verification)

	// creating tables
	err = DB.AutoMigrate(users, configs, runningInstances, verification)
	if err != nil {
		return err
	}
	return nil
}