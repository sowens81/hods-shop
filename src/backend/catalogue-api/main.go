package main

import (
	"catalogue-api/handlers"
	"catalogue-api/middleware"
	"catalogue-api/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize CosmosDB repository
	repo, err := services.NewCosmosRepository()
	if err != nil {
		log.Fatalf("Error initializing CosmosDB repository: %v", err)
	}

	// Initialize handlers
	catalogueHandler := handlers.NewCatalogueHandler(repo)

	// Initialize Gin router
	r := gin.Default()
	r.Use(middleware.LoggingMiddleware())

	// Routes
	r.GET("/catalogue", catalogueHandler.GetCatalogue)
	// r.GET("/catalogue/:id", catalogueHandler.GetCatalogueItem)
	// r.GET("/catalogue/size", catalogueHandler.GetCatalogueSize)
	// r.GET("/tags", catalogueHandler.GetTags)
	// r.POST("/catalogue", catalogueHandler.CreateCatalogueItem)
	// r.PUT("/catalogue/:id", catalogueHandler.UpdateCatalogueItem)
	// r.DELETE("/catalogue/:id", catalogueHandler.DeleteCatalogueItem)

	// Start the server
	fmt.Println("Server is running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
