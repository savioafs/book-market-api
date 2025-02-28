package entity

import (
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/utils"
)

type Client struct {
	ID    string `json:"id"`
	Name  string `json:"name"  gorm:"primaryKey"`
	Email string `json:"email"`
	Phone string `json:"phone" gorm:"uniqueIndex"`
}

func NewClient(name string, email string, phone string) (*Client, error) {
	id, err := utils.NewID()
	if err != nil {
		return nil, err
	}

	client := &Client{
		ID:    id,
		Name:  name,
		Email: email,
		Phone: phone,
	}

	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return common.ErrNameIsRequired
	}

	if c.Email == "" {
		return common.ErrEmailIsRequired
	}

	if c.Phone == "" {
		return common.ErrPhoneIsRequired
	}

	return nil
}
