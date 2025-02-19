package entity

import (
	"errors"
	"github.com/savioafs/book-market/internal/utils"
	"time"
)

var (
	ErrSaleIsRequired       = errors.New("sale is required")
	ErrRatingIsRequired     = errors.New("rating is required")
	ErrCommissionIsRequired = errors.New("commission is required")
)

type Review struct {
	ID        string    `json:"id" gorm:"primary_key"`
	SaleID    string    `json:"sale_id"`
	Sale      Sale      `json:"sale" gorm:"foreignKey:SaleID"`
	Rating    float32   `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

func NewReview(sale Sale, rating float32, comment string) (*Review, error) {
	id, err := utils.NewID()
	if err != nil {
		return nil, err
	}

	review := &Review{
		ID:        id,
		Sale:      sale,
		Rating:    rating,
		Comment:   comment,
		CreatedAt: time.Now(),
	}

	err = review.Validate()
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (r *Review) Validate() error {
	if r.Sale.ID == "" {
		return ErrSaleIsRequired
	}

	if r.Rating == 0.0 {
		return ErrRatingIsRequired
	}

	if r.Comment == "" {
		return ErrCommissionIsRequired
	}

	return nil
}
