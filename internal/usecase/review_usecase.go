package usecase

import (
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/converter"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
	"github.com/savioafs/book-market/internal/repository"
)

type ReviewUseCase struct {
	reviewRepository repository.ReviewStorer
	saleRepository   repository.SaleStorer
}

func NewReviewUseCase(reviewRepository repository.ReviewStorer, saleRepository repository.SaleStorer) *ReviewUseCase {
	return &ReviewUseCase{
		reviewRepository: reviewRepository,
		saleRepository:   saleRepository,
	}
}

func (u *ReviewUseCase) CreateReview(reviewInput dto.ReviewInputDTO) (dto.ReviewOutputDTO, error) {
	reviewExists, err := u.reviewRepository.ExistsReview(reviewInput.SaleID)
	if err != nil {
		return dto.ReviewOutputDTO{}, err
	}

	if reviewExists {
		return dto.ReviewOutputDTO{}, common.ErrSaleAlreadyReviewed
	}

	sale, err := u.saleRepository.GetSaleByID(reviewInput.SaleID)
	if err != nil {
		return dto.ReviewOutputDTO{}, err
	}

	if sale == nil {
		return dto.ReviewOutputDTO{}, err
	}

	review, err := entity.NewReview(
		*sale,
		reviewInput.Rating,
		reviewInput.Comment,
	)
	if err != nil {
		return dto.ReviewOutputDTO{}, err
	}

	reviewOutput := converter.ConvertReviewOutput(review, sale)

	return reviewOutput, nil
}
