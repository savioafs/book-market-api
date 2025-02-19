package service

import (
	"github.com/savioafs/book-market/internal/repository"
	"gorm.io/gorm"
	"log"
	"time"
)

func MonitorCouponsByExpirationDate(db *gorm.DB) {
	for {
		time.Sleep(time.Hour)

		discountCouponsRepository := repository.NewDiscountCouponRepositoryGorm(db)

		coupons, err := discountCouponsRepository.GetActiveDiscountsCoupons()
		if err != nil {
			log.Printf("fail to get active discount coupons: %s", err.Error())
			continue
		}

		for _, coupon := range *coupons {
			if coupon.ExpirationDate.Before(time.Now()) {
				err = discountCouponsRepository.DisableDiscountCoupon(coupon.ID)
				if err != nil {
					log.Printf("fail to discount coupon: %s", err.Error())
					continue
				}
			}
		}
	}
}
