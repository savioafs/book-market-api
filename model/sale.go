package model

import "time"

type Sale struct {
	ID             string         `json:"id"`
	Book           Book           `json:"book"`
	Seller         Seller         `json:"seller"`
	BuyerName      string         `json:"buyer_name"`
	Quantity       int            `json:"quantity"`
	TotalPrice     float64        `json:"total_price"`
	SaleDate       time.Time      `json:"sale_date"`
	DiscountCoupon DiscountCoupon `json:"discount_coupon"`
	IsReviewed     bool           `json:"isReviewed"`
}
