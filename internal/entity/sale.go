package entity

import (
	"github.com/savioafs/book-market/internal/utils"
	"time"
)

type Sale struct {
	ID               string         `json:"id" gorm:"primary_key"`
	Code             string         `json:"code"`
	Books            []Book         `json:"books" gorm:"many2many:sale_books"`
	SellerID         string         `json:"seller_id"`
	Seller           Seller         `json:"seller" gorm:"foreignKey:SellerID"`
	BuyerName        string         `json:"buyer_name"`
	Quantity         int            `json:"quantity"`
	TotalPrice       float64        `json:"total_price"`
	SaleDate         time.Time      `json:"sale_date"`
	DiscountCouponID string         `json:"discount_coupon_id"`
	DiscountCoupon   DiscountCoupon `json:"discount_coupon" gorm:"foreignKey:DiscountCouponID"`
	IsReviewed       bool           `json:"isReviewed"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
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
		SellerID:       seller.ID,
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
