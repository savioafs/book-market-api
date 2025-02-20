package usecase

import (
	"github.com/savioafs/book-market/internal/repository"
)

type DiscountCouponUseCase struct {
	repository repository.DiscountCouponStorer
}

func NewDiscountCouponUseCase(repository repository.DiscountCouponStorer) *DiscountCouponUseCase {
	return &DiscountCouponUseCase{repository: repository}
}

//func (u *DiscountCouponUseCase) CreateDiscountCoupon(couponInput *entity.DiscountCoupon) error {}
