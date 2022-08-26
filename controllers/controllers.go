package controllers

import (
	"net/http"

	"github.com/anhskrttt/todoapp-go-crud/initializers"
	"github.com/anhskrttt/todoapp-go-crud/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // H is a shortcut for map[string]interface{} (https://pkg.go.dev/github.com/gin-gonic/gin#H)
		"response": "pong",
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
	// Step 01. Get data off req body
	var taskData struct {
		Title       string
		Description string
		Status      bool
	}

	if err := c.ShouldBind(&taskData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Step 02. Add task to database
	task := models.Task{
		Title:       taskData.Title,
		Description: taskData.Description,
		Status:      taskData.Status,
	}
	result := initializers.DB.Create(&task)

	// Step 03. Check if there's any error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "Failed to create a new task",
		})
		return
	}

	// Step 04. Respond with the data added
	c.JSON(http.StatusOK, gin.H{
		"response": task,
	})
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

	// Should I check it? It runs fine even when I dont do this
	// Check if id is found
	var task models.Task
	resultFind := initializers.DB.First(&task, id)

	if resultFind.Error != nil {
		notFoundResponse(c, "Task not found")
		return
	}

	// Delete a task by id
	resultDelete := initializers.DB.Delete(&models.Task{}, id)

	if resultDelete.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "deleted",
	})
}

func DeleteAllTasks(c *gin.Context) {
	var tasks []models.Task
	resultFind := initializers.DB.Find(&tasks)

	if resultFind.Error != nil {
		c.Status(400)
		return
	}

	length := len(tasks)

	// Should I check this?
	// Observation through Postman: Faster ~50%?
	if length != 0 {
		// The following commands are the same
		// Command 01. initializers.DB.Exec("DELETE FROM tasks")
		// Question: Is this permanent deleting? As the id continuously raises.
		// f.e I deleted tasks from id 1 to 10
		// Now when I create new tasks, they will start at id 11 (not 1)
		// Command 02.
		resultDelete := initializers.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Task{})

		if resultDelete.Error != nil {
			c.Status(400)
			return
		}
	}

	successfulResponseWithLenParam(c, length)
}

// Utilities
func successfulResponseWithTaskModel(c *gin.Context, t *models.Task) {
	c.JSON(http.StatusOK, gin.H{
		"response": t,
	})
}

func notFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"response": message,
	})
}

func successfulResponseWithLenParam(c *gin.Context, length int) {
	c.JSON(http.StatusOK, gin.H{
		"response": "successful",
		"length":   length,
	})
}
