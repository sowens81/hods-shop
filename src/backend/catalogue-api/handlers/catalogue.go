package handlers

import (
	"catalogue-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CatalogueHandler is a struct that holds a reference to the CosmosDB repository.
type CatalogueHandler struct {
	repo *services.CosmosRepository
}

// NewCatalogueHandler creates a new handler instance with a reference to the repository.
func NewCatalogueHandler(repo *services.CosmosRepository) *CatalogueHandler {
	return &CatalogueHandler{repo: repo}
}

// GetCatalogue handles the GET request to retrieve all catalogue items.
func (h *CatalogueHandler) GetCatalogue(c *gin.Context) {
	ctx := c.Request.Context()
	items, err := h.repo.GetCatalogueItems(ctx)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError, "Failed to fetch catalogue items")
		return
	}
	c.JSON(http.StatusOK, items)
}

// // GetCatalogueItem handles the GET request to retrieve a catalogue item by its ID.
// func (h *CatalogueHandler) GetCatalogueItem(c *gin.Context) {
// 	id := c.Param("id")
// 	ctx := c.Request.Context()
// 	item, err := h.repo.GetCatalogueItemByID(ctx, id)
// 	if err != nil {
// 		handleError(c, err, http.StatusInternalServerError, "Failed to fetch item")
// 		return
// 	}
// 	if item == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, item)
// }

// // GetCatalogueSize handles the GET request to retrieve the catalogue size.
// func (h *CatalogueHandler) GetCatalogueSize(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	size, err := h.repo.GetCatalogueSize(ctx)
// 	if err != nil {
// 		handleError(c, err, http.StatusInternalServerError, "Failed to get catalogue size")
// 		return
// 	}
// 	c.JSON(http.StatusOK, dto.GetSizeResponse{Size: size})
// }

// // GetTags handles the GET request to retrieve tags.
// func (h *CatalogueHandler) GetTags(c *gin.Context) {
// 	// You can add a query here to fetch tags from CosmosDB or any other source
// 	tags := []string{"tag1", "tag2", "tag3"}
// 	c.JSON(http.StatusOK, dto.ListResponse3{Tags: tags})
// }

// // CreateCatalogueItem handles the POST request to create a new catalogue item.
// func (h *CatalogueHandler) CreateCatalogueItem(c *gin.Context) {
// 	var request dto.CreateItemRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
// 		return
// 	}

// 	// Handle the creation logic here, e.g., inserting the new item into CosmosDB
// 	ctx := c.Request.Context()
// 	item, err := h.repo.CreateItem(ctx, request)
// 	if err != nil {
// 		handleError(c, err, http.StatusInternalServerError, "Failed to create item")
// 		return
// 	}

// 	c.JSON(http.StatusCreated, item)
// }

// // UpdateCatalogueItem handles the PUT request to update an existing catalogue item by ID.
// func (h *CatalogueHandler) UpdateCatalogueItem(c *gin.Context) {
// 	id := c.Param("id")
// 	var request dto.UpdateItemRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
// 		return
// 	}

// 	// Handle the update logic here, e.g., updating the item in CosmosDB
// 	ctx := c.Request.Context()
// 	updatedItem, err := h.repo.UpdateItem(ctx, id, request)
// 	if err != nil {
// 		handleError(c, err, http.StatusInternalServerError, "Failed to update item")
// 		return
// 	}

// 	if updatedItem == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, updatedItem)
// }

// handleError is a utility function for structured error handling.
func handleError(c *gin.Context, err error, statusCode int, message string) {
	// Log the error (you can add more detailed logging for production)
	c.JSON(statusCode, gin.H{"error": message, "details": err.Error()})
}
