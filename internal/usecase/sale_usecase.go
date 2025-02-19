package usecase

import "github.com/savioafs/book-market/internal/repository"

type SaleUseCase struct {
	repository repository.SaleStorer
}

func NewSaleUseCase(repository repository.SaleStorer) *SaleUseCase {
	return &SaleUseCase{repository: repository}
}
