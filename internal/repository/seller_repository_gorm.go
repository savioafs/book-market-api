package repository

import (
	"database/sql"
	"errors"
	"github.com/savioafs/book-market/internal/entity"
	"gorm.io/gorm"
	"time"
)

type SellerRepositoryGorm struct {
	DB *sql.DB
}

func NewSellerRepositoryGorm(db *sql.DB) *SellerRepositoryGorm {
	return &SellerRepositoryGorm{DB: db}
}

func (r *SellerRepositoryGorm) CreateSeller(seller *entity.Seller) error {
	return r.DB.Create(seller).Error
}

func (r *SellerRepositoryGorm) GetAllSellers() ([]*entity.Seller, error) {
	var sellers []*entity.Seller

	err := r.DB.Find(&sellers).Error
	if err != nil {
		return nil, err
	}

	return sellers, nil
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (r *SellerRepositoryGorm) GetSellerByName(name string) (*entity.Seller, error) {
	var seller *entity.Seller

	err := r.DB.First(&seller, "name = ?", name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (r *SellerRepositoryGorm) GetSellerByPhone(phone string) (*entity.Seller, error) {
	var seller *entity.Seller

	err := r.DB.First(&seller, "phone = ?", phone).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

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

func (r *SellerRepositoryGorm) ExistsSeller(name, email, phone string) (bool, error) {
	var count int64

	err := r.DB.Model(&entity.Seller{}).Where("lower(name) LIKE lower(?) OR email = ? OR phone = ?  ", name, email, phone).Count(&count).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, nil
	}

	return count > 0, nil
}
