package entity

import (
	"errors"
	"github.com/savioafs/book-market/internal/utils"
	"strings"
	"time"
)

var (
	ErrCodeIsRequired               = errors.New("code is required")
	ErrDiscountPercentageIsRequired = errors.New("discount percentage is required")
	ErrExpirationDateIsRequired     = errors.New("expiration date is required")
	ErrUsageLimitIsRequired         = errors.New("usage limit is required")
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
		Active:             true,
		CreatedAt:          time.Now(),
	}

	err = discountCoupon.Validate()
	if err != nil {
		return nil, err
	}

	return discountCoupon, nil
}

func (d *DiscountCoupon) Validate() error {
	if d.Code == "" {
		return ErrCodeIsRequired
	}

	if d.DiscountPercentage <= 0 {
		return ErrDiscountPercentageIsRequired
	}

	if d.ExpirationDate.IsZero() {
		return ErrExpirationDateIsRequired
	}

	if d.UsageLimit <= 0 {
		return ErrUsageLimitIsRequired
	}

	return nil
}
