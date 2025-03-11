package dto

import (
	"time"
)

type SellerForSaleDTO struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type BookForSaleDTO struct {
	Title  string  `json:"name"`
	Author string  `json:"code"`
	Price  float64 `json:"price"`
}

type SaleInputDTO struct {
	BooksIDs         []string `json:"books_ids" `
	SellerID         string   `json:"seller_id"`
	ClientPhone      string   `json:"client_phone"`
	DiscountCouponID string   `json:"discount_coupon_id"`
}

type SaleOutputDTO struct {
	ID                 string                 `json:"id" gorm:"primary_key"`
	Code               string                 `json:"code"`
	Books              []BookForSaleDTO       `json:"books" `
	Seller             SellerForSaleDTO       `json:"seller"`
	Client             ClientOutputForSaleDTO `json:"client"`
	Quantity           int                    `json:"quantity"`
	TotalPrice         float64                `json:"total_price"`
	DiscountPercentage float64                `json:"discount_percentage"`
	DiscountCoupon     string                 `json:"discount_coupon"`
	FinalPrice         float64                `json:"final_price"`
	SaleDate           time.Time              `json:"sale_date"`
	IsReviewed         bool                   `json:"isReviewed"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
}
