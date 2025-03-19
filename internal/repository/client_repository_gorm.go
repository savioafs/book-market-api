package repository

import (
	"database/sql"
	"errors"
	"github.com/savioafs/book-market/internal/entity"
	"gorm.io/gorm"
)

type ClientRepositoryGorm struct {
	DB *sql.DB
}

func NewClientRepositoryGorm(db *sql.DB) *ClientRepositoryGorm {
	return &ClientRepositoryGorm{
		DB: db,
	}
}

func (r *ClientRepositoryGorm) CreateClient(client *entity.Client) error {
	return r.DB.Create(&client).Error
}

func (r *ClientRepositoryGorm) GetClientByPhone(phone string) (*entity.Client, error) {
	var client *entity.Client

	err := r.DB.First(&client, "phone = ?", phone).Error
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (r *ClientRepositoryGorm) UpdateClient(buyer *entity.Client) error {
	return nil
}

func (r *ClientRepositoryGorm) ExistsClient(name, email, phone string) (bool, error) {
	var count int64

	err := r.DB.Model(&entity.Client{}).Where("lower(name) LIKE lower(?) OR email = ? OR phone = ?  ", name, email, phone).Count(&count).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, nil
	}

	return count > 0, nil
}
