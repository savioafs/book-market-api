package converter

import (
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
)

func SaleToOutputDTO(sale *entity.Sale, discountCoupon *entity.DiscountCoupon, bookOutput []dto.BookForSaleDTO, sellerOutput dto.SellerForSaleDTO, clientOutput dto.ClientOutputForSaleDTO, discountPercentage, finalPrice float64) dto.SaleOutputDTO {
	return dto.SaleOutputDTO{
		ID:                 sale.ID,
		Code:               sale.Code,
		Books:              bookOutput,
		Seller:             sellerOutput,
		Client:             clientOutput,
		Quantity:           sale.Quantity,
		TotalPrice:         sale.TotalPrice,
		DiscountPercentage: discountPercentage,
		DiscountCoupon:     discountCoupon.Code,
		FinalPrice:         finalPrice,
		SaleDate:           sale.SaleDate,
		CreatedAt:          sale.CreatedAt,
	}
}
