package entity

import (
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/utils"
	"strings"
	"time"
)

type DiscountCoupon struct {
	ID                 string    `json:"id" gorm:"primary_key"`
	Code               string    `json:"code"`
	DiscountPercentage float64   `json:"discount_percent"`
	ExpirationDate     time.Time `json:"expiration_date"`
	UsageLimit         int       `json:"usage_limit"`
	UsedCount          int       `json:"used_count"`
	Active             bool      `json:"active"`
	CreatedAt          time.Time `json:"created_at"`
}

func NewDiscountCoupon(code string, discountPercentage float64, expirationDate time.Time, usageLimit int) (DiscountCoupon, error) {
	id, err := utils.NewID()
	if err != nil {
		return DiscountCoupon{}, err
	}

	discountCoupon := DiscountCoupon{
		ID:                 id,
		Code:               strings.ToUpper(code),
		DiscountPercentage: discountPercentage,
		ExpirationDate:     expirationDate,
		UsageLimit:         usageLimit,
		Active:             true,
		CreatedAt:          time.Now(),
	}

	err = discountCoupon.Validate()
	if err != nil {
		return DiscountCoupon{}, err
	}

	return discountCoupon, nil
}

func (d *DiscountCoupon) Validate() error {
	if d.Code == "" {
		return common.ErrCodeIsRequired
	}

	if d.DiscountPercentage <= 0 {
		return common.ErrDiscountPercentageIsRequired
	}

	if d.ExpirationDate.IsZero() {
		return common.ErrExpirationDateIsRequired
	}

	if d.UsageLimit <= 0 {
		return common.ErrUsageLimitIsRequired
	}

	return nil
}
