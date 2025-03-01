package repository

import (
	"github.com/savioafs/book-market/internal/entity"
	"gorm.io/gorm"
)

type ClientRepositoryGorm struct {
	DB *gorm.DB
}

func NewClientRepositoryGorm(db *gorm.DB) *ClientRepositoryGorm {
	return &ClientRepositoryGorm{
		DB: db,
	}
}

func (r *ClientRepositoryGorm) CreateClient(buyer *entity.Client) error {
	return nil
}

func (r *ClientRepositoryGorm) GetClientByPhone(phone string) (*entity.Client, error) {
	return nil, nil
}

func (r *ClientRepositoryGorm) UpdateClient(buyer *entity.Client) error {
	return nil
}
