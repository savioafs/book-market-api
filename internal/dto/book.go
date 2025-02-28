package dto

import (
	"time"
)

type BookInputDTO struct {
	Title         string  `json:"title"`
	ImageURL      string  `json:"image_url"`
	Author        string  `json:"author"`
	Publisher     string  `json:"publisher"`
	ISBN          string  `json:"isbn"`
	Price         float64 `json:"price"`
	Stock         int     `json:"stock"`
	Category      string  `json:"category"`
	PublishedYear int     `json:"published_year"`
	Description   string  `json:"description"`
}

type BookOutputDTO struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	ImageURL      string    `json:"image_url"`
	Author        string    `json:"author"`
	Publisher     string    `json:"publisher"`
	ISBN          string    `json:"isbn"`
	Price         float64   `json:"price"`
	Stock         int       `json:"stock"`
	Category      string    `json:"category"`
	PublishedYear int       `json:"published_year"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
