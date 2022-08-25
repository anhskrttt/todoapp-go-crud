package controllers

import (
	"net/http"

	"github.com/anhskrttt/todoapp-go-crud/initializers"
	"github.com/anhskrttt/todoapp-go-crud/models"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // H is a shortcut for map[string]interface{} (https://pkg.go.dev/github.com/gin-gonic/gin#H)
		"message": "pong",
	})
}

func GetAllTasks(c *gin.Context) {
	var tasks []models.Task
	result := initializers.DB.Find(&tasks)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"length": len(tasks),
		"task":   tasks,
	})
}

func GetTaskById(c *gin.Context) {
	// Get id off req body
	id := c.Param("id")

	// Find task with id
	var task models.Task
	result := initializers.DB.First(&task, id)

	if result.Error != nil {
		notFoundResponse(c, "Task not found")
		return
	}

	// Respond
	successfulResponseWithTaskModel(c, &task)
}

func CreateTask(c *gin.Context) {
	// Get data off req body
	var taskData struct {
		Title       string
		Description string
		Status      bool
	}

	c.Bind(&taskData)

	task := models.Task{
		Title:       taskData.Title,
		Description: taskData.Description,
		Status:      taskData.Status,
	}

	// Add task to database
	result := initializers.DB.Create(&task)

	if result.Error != nil {
		c.Status(400)
		return
	}

	successfulResponseWithTaskModel(c, &task)
}

func UpdateTaskById(c *gin.Context) {
	// Get id off req body
	id := c.Param("id")

	// Get data off req body
	var taskData struct {
		Title       string
		Description string
		Status      bool
	}

	c.Bind(&taskData)

	// Find task by id to update
	var task models.Task
	result := initializers.DB.First(&task, id)

	if result.Error != nil {
		notFoundResponse(c, "Task not found")
		return
	}

	// Update content to task element found
	initializers.DB.Model(&task).Updates(models.Task{
		Title:       taskData.Title,
		Description: taskData.Description,
		Status:      taskData.Status,
	})

	successfulResponseWithTaskModel(c, &task)
}

func DeleteTaskById(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Task{}, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}

// Utilities
func successfulResponseWithTaskModel(c *gin.Context, t *models.Task) {
	c.JSON(http.StatusOK, gin.H{
		"response": t,
	})
}

func notFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": message,
	})
}
