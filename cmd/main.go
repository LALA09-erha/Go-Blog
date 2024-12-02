package main

import (
	"Go-Blog/config"
	"Go-Blog/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDB()

	// Setup Gin router
	r := gin.Default()

	// Initialize routes
	routes.InitRoutes(r)

	routes.UserRoutes(r)


	// Start the server
	log.Fatal(r.Run(":8080"))
}
