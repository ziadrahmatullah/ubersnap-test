package main

import (
	"os"
	"ubersnap-test/logger"
	"ubersnap-test/migration"
	"ubersnap-test/repository"
)

func main() {
	logger.SetLogrusLogger()

	_ = os.Setenv("APP_ENV", "debug")

	db, err := repository.GetConnection()
	if err != nil {
		logger.Log.Error(err)
	}

	migration.Seed(db)
}
