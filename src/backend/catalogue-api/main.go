package main

import (
	"catalogue-api/cosmosdb"
	"catalogue-api/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize CosmosDB repository
	repo, err := cosmosdb.NewCosmosRepository()
	if err != nil {
		log.Fatalf("Error initializing CosmosDB repository: %v", err)
	}

	// Initialize handlers
	catalogueHandler := handlers.NewCatalogueHandler(repo)

	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.GET("/catalogue", catalogueHandler.GetCatalogue)
	r.GET("/catalogue/:id", catalogueHandler.GetCatalogueItem)
	r.GET("/catalogue/size", catalogueHandler.GetCatalogueSize)
	r.GET("/tags", catalogueHandler.GetTags)
	r.POST("/catalogue", catalogueHandler.CreateCatalogueItem)
	r.PUT("/catalogue/:id", catalogueHandler.UpdateCatalogueItem)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
