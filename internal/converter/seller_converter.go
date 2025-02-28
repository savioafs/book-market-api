package converter

import (
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
)

func SellerToOutputSaleDTO(seller *entity.Seller) dto.SellerForSaleDTO {
	return dto.SellerForSaleDTO{
		Name: seller.Name,
		Code: seller.Code,
	}
}
