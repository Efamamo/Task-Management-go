package router

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter() {
	r := gin.Default()
	r.GET("/tasks")
	r.GET("tasks/:id")
	r.PUT("/tasks/:id")
	r.DELETE("/tasks/:id")
	r.POST("/tasks")

}
