package converters

import (
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
)

func ConvertClientOutput(client *entity.Client) dto.ClientOutputDTO {
	return dto.ClientOutputDTO{
		Name:  client.Name,
		Email: client.Email,
		Phone: client.Phone,
	}
}

func ConvertSellerOutput(seller *entity.Seller) dto.SellerForSaleDTO {
	return dto.SellerForSaleDTO{
		Name: seller.Name,
		Code: seller.Code,
	}
}

func ConvertBookOutput(book *entity.Book) dto.BookForSaleDTO {
	return dto.BookForSaleDTO{
		Title:  book.Title,
		Author: book.Author,
		Price:  book.Price,
	}
}

func ConvertSaleOutput(sale *entity.Sale, discountCoupon *entity.DiscountCoupon, bookOutput []dto.BookForSaleDTO, sellerOutput dto.SellerForSaleDTO, clientOutput dto.ClientOutputDTO, discountPercentage, finalPrice float64) dto.SaleOutputDTO {
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
