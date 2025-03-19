package entity

import (
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/utils"
	"time"
)

// TODO: remember to total price return
// TODO: implement apply discount coupon
type Sale struct {
	ID           string    `json:"id"`
	Code         string    `json:"code"`
	Books        []string  `json:"books"`
	SellerCode   string    `json:"seller_code"`
	ClientPhone  string    `json:"client_phone"`
	Quantity     int       `json:"quantity"`
	DiscountCode string    `json:"discount_coupon"`
	IsReviewed   bool      `json:"is_reviewed"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewSale(books []string, sellerCode string, clientPhone string, discountCode string) (*Sale, error) {
	id, err := utils.NewID()
	if err != nil {
		return nil, err
	}

	code, err := utils.NewCode(sellerCode, 12)
	if err != nil {
		return nil, err
	}

	sale := &Sale{
		ID:           id,
		Code:         code,
		Books:        books,
		SellerCode:   sellerCode,
		ClientPhone:  clientPhone,
		Quantity:     len(books),
		DiscountCode: discountCode,
		IsReviewed:   false,
		CreatedAt:    time.Now(),
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

	if s.SellerCode == "" {
		return common.ErrSellerCodeIsRequired
	}

	if s.ClientPhone == "" {
		return common.ErrClientPhoneIsRequired
	}

	return nil
}

/*
func (s *Sale) ApplyDiscountPercentageAndFinalPrice(discountCoupon *DiscountCoupon) (float64, float64, error) {
	var discountPercentage float64
	var finalPrice float64

	if discountCoupon.ID == "" {
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
*/
