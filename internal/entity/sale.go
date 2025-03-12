package entity

import (
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/utils"
	"time"
)

type Sale struct {
	ID               string         `json:"id" gorm:"primary_key"`
	Code             string         `json:"code"`
	Books            []Book         `json:"books" gorm:"many2many:sale_books"`
	SellerID         string         `json:"seller_id"`
	Seller           Seller         `json:"seller" gorm:"foreignKey:SellerID"`
	ClientPhone      string         `json:"client_phone"`
	Client           Client         `json:"client" gorm:"foreignKey:ClientPhone;references:Phone"`
	Quantity         int            `json:"quantity"`
	TotalPrice       float64        `json:"total_price"`
	SaleDate         time.Time      `json:"sale_date"`
	DiscountCouponID string         `json:"discount_coupon_id"`
	DiscountCoupon   DiscountCoupon `json:"discount_coupon" gorm:"foreignKey:DiscountCouponID"`
	IsReviewed       bool           `json:"is_reviewed"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

func NewSale(books []Book, seller Seller, client Client, discountCoupon DiscountCoupon) (*Sale, error) {
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
		Client:         client,
		Quantity:       len(books),
		TotalPrice:     totalPrice,
		SaleDate:       time.Now(),
		DiscountCoupon: discountCoupon,
		IsReviewed:     false,
		CreatedAt:      time.Now(),
	}

	err = sale.Validate()
	if err != nil {
		return nil, err
	}

	return sale, nil
}

func (s *Sale) Validate() error {

	if s.Books == nil {
		return common.ErrBookIsRequired
	}

	if s.Seller.ID == "" {
		return common.ErrSellerIsRequired
	}

	if s.Client.Name == "" {
		return common.ErrBuyerIsRequired
	}

	return nil
}

func (s *Sale) ApplyDiscountPercentageAndFinalPrice(discountCoupon *DiscountCoupon) (float64, float64, error) {
	var discountPercentage float64
	var finalPrice float64

	if discountCoupon == nil {
		discountPercentage = 0
		finalPrice = s.TotalPrice
		return discountPercentage, finalPrice, nil
	}

	if !discountCoupon.Active {
		return discountPercentage, finalPrice, common.ErrDiscountCouponIsNotActive
	}

	discountPercentage = discountCoupon.DiscountPercentage
	finalPrice = s.TotalPrice * (1 - (discountPercentage / 100))

	return discountPercentage, finalPrice, nil
}
