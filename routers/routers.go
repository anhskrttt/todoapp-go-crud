package routers

import (
	"net/http"

	"github.com/anhskrttt/todoapp-go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func Routers() http.Handler {
	// Default With the Logger and Recovery middleware already attached
	r := gin.Default()

	// More about gin Context: https://stackoverflow.com/questions/63522977/how-does-context-struct-work-in-golang-gin-framework
	// The gin Context is a structure that contains both the http.Request and the http.Response
	// Test API
	r.GET("/ping", controllers.Ping)

	// My API
	r.GET("/api/tasks", controllers.GetAllTasks)
	r.GET("/api/tasks/:id", controllers.GetTaskById)

	r.POST("/api/task", controllers.CreateTask)

	r.PUT("/api/tasks/:id", controllers.UpdateTaskById)
	
	r.DELETE("/api/tasks/:id", controllers.DeleteTaskById)

	return r
}
