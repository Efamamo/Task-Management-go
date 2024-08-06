package router

import (
	controller "github.com/Task-Management-go/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() {
	r := gin.Default()
	r.GET("/tasks", controller.GetTasks)
	r.GET("tasks/:id", controller.GetTaskById)
	r.PUT("/tasks/:id", controller.UpdateItem)
	r.DELETE("/tasks/:id")
	r.POST("/tasks")

	r.Run()

}
