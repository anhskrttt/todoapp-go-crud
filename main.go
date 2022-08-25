package main

import (
	"log"
	"net/http"
	"os"

	"github.com/anhskrttt/todoapp-go-crud/initializers"
	"github.com/anhskrttt/todoapp-go-crud/routers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := routers.Routers()

	// More about gin Context: https://stackoverflow.com/questions/63522977/how-does-context-struct-work-in-golang-gin-framework
	// The gin Context is a structure that contains both the http.Request and the http.Response

	err := http.ListenAndServe(os.Getenv("PORT"), r)

	if err != nil {
		log.Fatal("Failed to start server....")
	}
}
