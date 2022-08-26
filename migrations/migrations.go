package main

import (
	"github.com/anhskrttt/todoapp-go-crud/initializers"
	"github.com/anhskrttt/todoapp-go-crud/models"
)

func init() {
	initializers.LoadEnv() // Not use yet
	initializers.ConnectDB()
}

func main() {
	// Create a new table
	initializers.DB.AutoMigrate(&models.Task{})

	// Drop a table
	// initializers.DB.Migrator().DropTable(&models.Task{})
}
