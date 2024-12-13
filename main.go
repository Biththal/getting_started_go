package main

import (
	"todo/config"
	"todo/models"
	"todo/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()

	// Migrate the database
	models.MigrateTodo()

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":8084")
}
