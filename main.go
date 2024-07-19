package main

import (
	"github.com/gin-gonic/gin"

	"todo/config"
	"todo/database"
	"todo/routes"
)

func main() {
	config.InitEnvConfig()
	database.InitDB()

	server := gin.Default()
	routes.Routes(server)
	server.Run(":8080")
	defer database.DB.Close()
}
