package repository

import (
	"errors"
	"github.com/savioafs/book-market/internal/entity"
	"gorm.io/gorm"
	"time"
)

type SellerRepositoryGorm struct {
	DB *gorm.DB
}

func NewSellerRepositoryGorm(db *gorm.DB) *SellerRepositoryGorm {
	return &SellerRepositoryGorm{DB: db}
}

func (r *SellerRepositoryGorm) CreateSeller(seller *entity.Seller) error {
	return r.DB.Create(seller).Error
}

func (r *SellerRepositoryGorm) GetSellerByID(id string) (*entity.Seller, error) {
	var seller *entity.Seller

	err := r.DB.First(&seller, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (r *SellerRepositoryGorm) GetSellerByEmail(email string) (*entity.Seller, error) {
	var seller *entity.Seller

	err := r.DB.First(&seller, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (r *SellerRepositoryGorm) UpdateSeller(seller *entity.Seller) error {
	err := r.DB.Model(&entity.Seller{}).Where("id = ?", seller.ID).Updates(map[string]interface{}{
		"name":       seller.Name,
		"email":      seller.Email,
		"phone":      seller.Phone,
		"updated_at": time.Now(),
	}).Error

	return err
}

func (r *SellerRepositoryGorm) DeleteSeller(id string) error {
	res := r.DB.Delete(&entity.Seller{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("seller not found")
	}

	return nil
}
