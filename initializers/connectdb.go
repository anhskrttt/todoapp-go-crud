package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// An ORM (object-relational mapper) library is a completely ordinary library written in your language of choice
	// that encapsulates the code needed to manipulate the data,
	// so you don't use SQL anymore; you interact directly with an object in the same language you're using.

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// When to use log.Fatal: https://stackoverflow.com/questions/33885235/should-a-go-package-ever-use-log-fatal-and-when
		log.Fatal("Failed to connect to the database.")
	}
}
