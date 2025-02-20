package dto

import (
	"time"
)

type ReviewInputDTO struct {
	SaleID  string  `json:"sale_id"`
	Rating  float32 `json:"rating"`
	Comment string  `json:"comment"`
}

type ReviewOutputDTO struct {
	ID        string    `json:"id"  `
	SaleID    string    `json:"sale_id"`
	Seller    string    `json:"seller"`
	SaleDate  time.Time `json:"sale_date"`
	Rating    float32   `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
