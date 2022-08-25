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
	initializers.DB.AutoMigrate(&models.Task{})
	// initializers.DB.Migrator().DropTable(&models.Task{})
}
