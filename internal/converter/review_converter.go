package converter

import (
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
)

func ConvertReviewOutput(review *entity.Review, sale *entity.Sale) dto.ReviewOutputDTO {
	return dto.ReviewOutputDTO{
		ID:        review.ID,
		SaleID:    sale.ID,
		Seller:    sale.Seller.Name,
		SaleDate:  sale.CreatedAt,
		Rating:    review.Rating,
		Comment:   review.Comment,
		CreatedAt: review.CreatedAt,
	}
}
