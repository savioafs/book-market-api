package entity

import (
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/utils"
	"time"
)

type Seller struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func NewSeller(name, email, phone string) (*Seller, error) {
	id, err := utils.NewID()
	if err != nil {
		return nil, err
	}

	code, err := utils.NewCode(name, 6)
	if err != nil {
		return nil, err
	}

	seller := &Seller{
		ID:        id,
		Code:      code,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Rating:    5,
		CreatedAt: time.Now(),
	}

	err = seller.Validate()
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (s *Seller) Validate() error {
	if s.Name == "" {
		return common.ErrNameIsRequired
	}

	if s.Email == "" {
		return common.ErrEmailIsRequired
	}

	if s.Phone == "" {
		return common.ErrPhoneIsRequired
	}

	return nil
}
