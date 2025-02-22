package main

import (
	"UnUpset/config"
	"UnUpset/models"
	"UnUpset/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	err := models.Migrate(config.DB)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	r := gin.Default()
	routes.SetupRoutesUsers(r)
	routes.SetupRoutesTodos(r)
	routes.SetupRoutesStatistics(r)
	log.Println("Server is running on http://localhost:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
