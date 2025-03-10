package converter

import (
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
)

func SellerToOutputDTO(seller *entity.Seller) dto.SellerOutputDTO {
	return dto.SellerOutputDTO{
		ID:        seller.ID,
		Code:      seller.Code,
		Name:      seller.Name,
		Email:     seller.Email,
		Phone:     seller.Phone,
		Rating:    seller.Rating,
		CreatedAt: seller.CreatedAt,
	}
}

func SellerToOutputSaleDTO(seller *entity.Seller) dto.SellerForSaleDTO {
	return dto.SellerForSaleDTO{
		Name: seller.Name,
		Code: seller.Code,
	}
}
