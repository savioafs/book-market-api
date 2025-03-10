package dto

import "time"

type SellerInputDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type SellerOutputDTO struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
