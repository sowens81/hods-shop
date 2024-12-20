package dto

type UpdateItemRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	ImageUrl    []string `json:"imageUrl"`
	Price       *float64 `json:"price"`
	Count       *int32   `json:"count"`
	Tag         []string `json:"tag"`
}
