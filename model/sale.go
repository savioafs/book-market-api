package model

import (
	"github.com/savioafs/book-market/utils"
	"time"
)

type Sale struct {
	ID             string         `json:"id"`
	Code           string         `json:"code"`
	Book           Book           `json:"book"`
	Seller         Seller         `json:"seller"`
	BuyerName      string         `json:"buyer_name"`
	Quantity       int            `json:"quantity"`
	TotalPrice     float64        `json:"total_price"`
	SaleDate       time.Time      `json:"sale_date"`
	DiscountCoupon DiscountCoupon `json:"discount_coupon"`
	IsReviewed     bool           `json:"isReviewed"`
	CreatedAt      time.Time      `json:"created_at"`
}

func NewSale(book Book, seller Seller, buyerName string, quantity int, totalPrice float64, discountCoupon DiscountCoupon) (*Sale, error) {
	id, err := utils.NewID()
	if err != nil {
		return nil, err
	}

	code, err := utils.NewCode(book.Title)
	if err != nil {
		return nil, err
	}

	sale := &Sale{
		ID:             id,
		Code:           code,
		Book:           book,
		Seller:         seller,
		BuyerName:      buyerName,
		Quantity:       quantity,
		TotalPrice:     totalPrice,
		SaleDate:       time.Now(),
		DiscountCoupon: discountCoupon,
		IsReviewed:     false,
		CreatedAt:      time.Now(),
	}

	return sale, nil
}
