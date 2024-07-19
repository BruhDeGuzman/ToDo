package routes

import "github.com/gin-gonic/gin"

func Routes(server *gin.Engine) {
	server.GET("/tasks", getAllTasksHandler)
	server.GET("/tasks/:id", getTaskHandler)
	server.POST("/tasks", createTaskHandler)
	server.PUT("/tasks/:id", updateTaskHandler)
	server.DELETE("tasks/:id", deleteTaskHandler)
}
