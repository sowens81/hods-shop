package models

type Image struct {
	ImageAltText string `json:"imageAltText" db:"image_alt_text"`
	ImageURL     string `json:"imageUrl" db:"image_url"`
}
