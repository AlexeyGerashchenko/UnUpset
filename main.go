package main

import (
	"log"
	"net/http"

	"UnUpset/config"
	"UnUpset/models"
	"UnUpset/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	err := models.Migrate(config.DB)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	r := gin.Default()
	routes.SetupRotesUsers(r)
	routes.SetupRotesTodos(r)
	routes.SetupRotesStatistics(r)
	routes.SetupRotesTimer(r)
	log.Println("Server is running on http://localhost:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}