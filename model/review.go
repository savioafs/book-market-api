package model

import (
	"github.com/savioafs/book-market/utils"
	"time"
)

type Review struct {
	ID        string    `json:"id"`
	Sale      Sale      `json:"sale"`
	Seller    Seller    `json:"seller"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

func NewReview(sale Sale, seller Seller, rating int, comment string) (*Review, error) {
	id, err := utils.NewID()
	if err != nil {
		return nil, err
	}

	review := &Review{
		ID:        id,
		Sale:      sale,
		Seller:    seller,
		Rating:    rating,
		Comment:   comment,
		CreatedAt: time.Now(),
	}

	return review, nil
}
