package main

import (
	"log"

	"Proyecto2/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database

	// Create a new Gin router
	router := gin.Default()

	// Create a new API instance
	api := api.API{
		Router: router,
	}
	// Configure routes
	api.SetupRoutes()

	// Start the server
	log.Println("Server started on http://localhost:8080")
	router.Run(":8080")
}
