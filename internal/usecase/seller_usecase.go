package usecase

import (
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/repository"
)

type SellerUseCase struct {
	repository repository.SellerStorer
}

func NewSellerUseCase(repository repository.SellerStorer) *SellerUseCase {
	return &SellerUseCase{repository: repository}
}

func (u *SellerUseCase) CreateSeller(sellerInput dto.SellerInputDTO) (dto.SellerOutputDTO, error) {

}
