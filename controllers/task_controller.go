package controller

import (
	"net/http"

	"github.com/Task-Management-go/data"
	model "github.com/Task-Management-go/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := data.GetTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task := data.GetTaskById(id)

	if task == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, *task)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedTask model.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := data.UpdateItem(id, updatedTask)

	if task == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, *task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task := data.DeleteTask(id)

	if task == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, *task)
}
