package entity

import (
	"github.com/savioafs/book-market/internal/utils"
	"gorm.io/gorm"
	"time"
)

type Sale struct {
	gorm.Model
	ID             string         `json:"id"`
	Code           string         `json:"code"`
	Books          []Book         `json:"books"`
	Seller         Seller         `json:"seller"`
	BuyerName      string         `json:"buyer_name"`
	Quantity       int            `json:"quantity"`
	TotalPrice     float64        `json:"total_price"`
	SaleDate       time.Time      `json:"sale_date"`
	DiscountCoupon DiscountCoupon `json:"discount_coupon"`
	IsReviewed     bool           `json:"isReviewed"`
	CreatedAt      time.Time      `json:"created_at"`
}

func NewSale(books []Book, seller Seller, buyerName string, discountCoupon DiscountCoupon) (*Sale, error) {
	id, err := utils.NewID()
	if err != nil {
		return nil, err
	}

	code, err := utils.NewCode(seller.Code, 12)
	if err != nil {
		return nil, err
	}

	totalPrice := 0.0

	for _, book := range books {
		totalPrice += book.Price
	}

	sale := &Sale{
		ID:             id,
		Code:           code,
		Books:          books,
		Seller:         seller,
		BuyerName:      buyerName,
		Quantity:       len(books),
		TotalPrice:     totalPrice,
		SaleDate:       time.Now(),
		DiscountCoupon: discountCoupon,
		IsReviewed:     false,
		CreatedAt:      time.Now(),
	}

	return sale, nil
}
