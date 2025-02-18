package entity

import (
	"github.com/savioafs/book-market/internal/utils"
	"gorm.io/gorm"
	"time"
)

type Seller struct {
	gorm.Model
	ID        string    `json:"id"`
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

	return seller, nil
}
