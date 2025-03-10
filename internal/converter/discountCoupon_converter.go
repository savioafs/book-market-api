package converter

import (
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
)

func DiscountCouponToOutputDTO(d *entity.DiscountCoupon) dto.DiscountCouponOutputDTO {
	return dto.DiscountCouponOutputDTO{
		ID:                 d.ID,
		Code:               d.Code,
		DiscountPercentage: d.DiscountPercentage,
		ExpirationDate:     d.ExpirationDate,
		UsageLimit:         d.UsageLimit,
		Active:             d.Active,
		CreatedAt:          d.CreatedAt,
	}
}
