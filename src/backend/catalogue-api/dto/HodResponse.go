package dto

import (
	"catalogue-api/models"
)

type HodResponse struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      []models.Image `json:"images"`
	Price       float64        `json:"price"`
	Count       int32          `json:"count"`
	Tag         []string       `json:"tag"`
}
