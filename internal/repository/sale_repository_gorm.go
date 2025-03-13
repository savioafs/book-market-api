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

func (r *SaleRepositoryGorm) ExistsRecentSale(sellerID, clientPhone string, booksIDs []string, withInMinutes int) (bool, error) {
	var count int64

	minTime := time.Now().Add(-time.Duration(withInMinutes) * time.Minute)

	err := r.DB.Table("sales").Joins("JOIN sale_books sb ON sb.sale_id = sales.id").
		Where("sales.seller_id = ? AND sales.client_phone = ? AND sales.created_at >= ?", sellerID, clientPhone, minTime).
		Where("sb.book_id IN ?", booksIDs).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
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
