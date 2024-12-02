package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ASCE-D/makers-suite/internal/handlers"
	"github.com/ASCE-D/makers-suite/pkg/database"
)

func main() {
	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize router
	r := gin.Default()

	// Setup routes
	handlers.SetupRoutes(r, db)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}