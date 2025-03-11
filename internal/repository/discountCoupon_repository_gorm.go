package repository

import (
	"errors"
	"github.com/savioafs/book-market/internal/entity"
	"gorm.io/gorm"
)

type DiscountCouponRepositoryGorm struct {
	DB *gorm.DB
}

func NewDiscountCouponRepositoryGorm(db *gorm.DB) *DiscountCouponRepositoryGorm {
	return &DiscountCouponRepositoryGorm{DB: db}
}

func (r *DiscountCouponRepositoryGorm) CreateDiscountCoupon(coupon *entity.DiscountCoupon) error {
	return r.DB.Create(coupon).Error
}

func (r *DiscountCouponRepositoryGorm) GetAllDiscountCoupons() ([]*entity.DiscountCoupon, error) {
	var discountCoupons []*entity.DiscountCoupon

	err := r.DB.Find(&discountCoupons).Error
	if err != nil {
		return nil, err
	}

	return discountCoupons, nil
}

func (r *DiscountCouponRepositoryGorm) CountUse(id string) error {
	var discountCoupon *entity.DiscountCoupon

	err := r.DB.First(&discountCoupon, "id = ?", id).Error
	if err != nil {
		return err
	}

	discountCoupon.UsedCount += 1

	return r.DB.Save(&discountCoupon).Error
}

func (r *DiscountCouponRepositoryGorm) GetDiscountCoupon(id string) (*entity.DiscountCoupon, error) {
	var discountCoupon *entity.DiscountCoupon

	err := r.DB.First(&discountCoupon, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return discountCoupon, nil
}

func (r *DiscountCouponRepositoryGorm) GetActiveDiscountsCoupons() (*[]entity.DiscountCoupon, error) {
	var discountCoupons []entity.DiscountCoupon

	err := r.DB.Where("active = ?", true).Find(&discountCoupons).Error
	if err != nil {
		return nil, err
	}

	return &discountCoupons, nil
}

func (r *DiscountCouponRepositoryGorm) DisableDiscountCoupon(id string) error {
	var discountCoupon entity.DiscountCoupon

	err := r.DB.First(&discountCoupon, "id = ?", id).Error
	if err != nil {
		return err
	}

	discountCoupon.Active = false

	return r.DB.Save(&discountCoupon).Error
}

func (r *DiscountCouponRepositoryGorm) ExistsDiscountCoupon(code string) (bool, error) {
	var count int64

	err := r.DB.Model(&entity.DiscountCoupon{}).Where("lower(code) = lower(?)", code).Count(&count).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
