package usecase

import (
	"github.com/savioafs/book-market/internal/converter"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
	"github.com/savioafs/book-market/internal/repository"
)

type DiscountCouponUseCase struct {
	repository repository.DiscountCouponStorer
}

func NewDiscountCouponUseCase(repository repository.DiscountCouponStorer) *DiscountCouponUseCase {
	return &DiscountCouponUseCase{repository: repository}
}

func (u *DiscountCouponUseCase) CreateDiscountCoupon(couponInput *entity.DiscountCoupon) (dto.DiscountCouponOutputDTO, error) {
	discountCoupon, err := entity.NewDiscountCoupon(
		couponInput.Code,
		couponInput.DiscountPercentage,
		couponInput.ExpirationDate,
		couponInput.UsageLimit,
		couponInput.UsedCount,
	)

	if err != nil {
		return dto.DiscountCouponOutputDTO{}, err
	}

	err = u.repository.CreateDiscountCoupon(discountCoupon)
	if err != nil {
		return dto.DiscountCouponOutputDTO{}, err
	}

	discountCouponOutput := converter.DiscountCouponToOutputDTO(discountCoupon)

	return discountCouponOutput, nil
}
