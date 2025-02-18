package entity

import (
	"github.com/savioafs/book-market/utils"
	"strings"
	"time"
)

type DiscountCoupon struct {
	ID                 string    `json:"id"`
	Code               string    `json:"code"`
	DiscountPercentage float64   `json:"discount_percent"`
	ExpirationDate     time.Time `json:"expiration_date"`
	UsageLimit         int       `json:"usage_limit"`
	UsedCount          int       `json:"used_count"`
	CreatedAt          time.Time `json:"created_at"`
}

func NewDiscountCoupon(code string, discountPercentage float64, expirationDate time.Time, usageLimit int, usedCount int) (*DiscountCoupon, error) {
	id, err := utils.NewID()
	if err != nil {
		return nil, err
	}

	discountCoupon := &DiscountCoupon{
		ID:                 id,
		Code:               strings.ToUpper(code),
		DiscountPercentage: discountPercentage,
		ExpirationDate:     expirationDate,
		UsageLimit:         usageLimit,
		UsedCount:          usedCount,
		CreatedAt:          time.Now(),
	}

	return discountCoupon, nil
}
