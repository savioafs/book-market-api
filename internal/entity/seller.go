package entity

import (
	"errors"
	"github.com/savioafs/book-market/internal/utils"
	"time"
)

var (
	ErrNameIsRequired  = errors.New("name is required")
	ErrEmailIsRequired = errors.New("email is required")
	ErrPhoneIsRequired = errors.New("phone is required")
)

type Seller struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time `json:"create_at"`
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
		return ErrNameIsRequired
	}

	if s.Email == "" {
		return ErrEmailIsRequired
	}

	if s.Phone == "" {
		return ErrPhoneIsRequired
	}
	
	return nil
}
