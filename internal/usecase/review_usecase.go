package usecase

import "github.com/savioafs/book-market/internal/repository"

type ReviewUseCase struct {
	repository repository.ReviewStorer
}

func NewReviewUseCase(repository repository.ReviewStorer) *ReviewUseCase {
	return &ReviewUseCase{repository: repository}
}
