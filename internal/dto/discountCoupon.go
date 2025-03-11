package dto

import "time"

type DiscountCouponInputDTO struct {
	Code               string    `json:"code"`
	DiscountPercentage float64   `json:"discount_percent"`
	ExpirationDate     time.Time `json:"expiration_date"`
	UsageLimit         int       `json:"usage_limit"`
	UsedCount          int       `json:"used_count"`
}

type DiscountCouponOutputDTO struct {
	ID                 string    `json:"id" gorm:"primary_key"`
	Code               string    `json:"code"`
	DiscountPercentage float64   `json:"discount_percent"`
	ExpirationDate     time.Time `json:"expiration_date"`
	UsageLimit         int       `json:"usage_limit"`
	UsedCount          int       `json:"used_count"`
	Active             bool      `json:"active"`
	CreatedAt          time.Time `json:"created_at"`
}
