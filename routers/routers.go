package routers

import (
	"net/http"

	"github.com/anhskrttt/todoapp-go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func Routers() http.Handler {
	// Default With the Logger and Recovery middleware already attached
	r := gin.Default()

	// Load CSS file
	r.Static("public", "public")

	// Load HTML file
	r.LoadHTMLGlob("templates/*")

	// Serve static files
	r.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"status": http.StatusOK,
			},
		)
	})

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
	r.DELETE("/api/tasks", controllers.DeleteAllTasks)

	return r
}
