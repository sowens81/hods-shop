package dto

type CreateItemRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	ImageUrl    []string `json:"imageUrl" binding:"required"`
	Price       float64  `json:"price" binding:"required"`
	Count       int32    `json:"count" binding:"required"`
	Tag         []string `json:"tag" binding:"required"`
}
