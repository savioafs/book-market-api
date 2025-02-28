package repository

import (
	"errors"
	"github.com/savioafs/book-market/internal/entity"
	"gorm.io/gorm"
	"time"
)

type SaleRepositoryGorm struct {
	DB *gorm.DB
}

func NewSaleRepositoryGorm(db *gorm.DB) *SaleRepositoryGorm {
	return &SaleRepositoryGorm{DB: db}
}

func (r *SaleRepositoryGorm) CreateSale(sale *entity.Sale) error {
	return r.DB.Create(sale).Error
}

func (r *SaleRepositoryGorm) GetAllSales() ([]*entity.Sale, error) {
	var sales []*entity.Sale

	err := r.DB.Find(&sales).Error
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (r *SaleRepositoryGorm) GetSaleByID(id string) (*entity.Sale, error) {
	var sale *entity.Sale

	err := r.DB.First(&sale, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return sale, nil
}

func (r *SaleRepositoryGorm) UpdateSale(sale *entity.Sale) error {
	err := r.DB.Model(&entity.Seller{}).Where("id = ?", sale.ID).Updates(map[string]interface{}{
		"books":       sale.Books,
		"client":      sale.Client,
		"quantity":    sale.Quantity,
		"total_price": sale.TotalPrice,
		"updated_at":  time.Now(),
	}).Error

	return err
}

func (r *SaleRepositoryGorm) DeleteSale(id string) error {
	res := r.DB.Delete(&entity.Sale{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("sale not found")
	}

	return nil
}
