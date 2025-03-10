package usecase

import (
	"github.com/savioafs/book-market/internal/converter"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
	"github.com/savioafs/book-market/internal/repository"
)

type SellerUseCase struct {
	repository repository.SellerStorer
}

func NewSellerUseCase(repository repository.SellerStorer) *SellerUseCase {
	return &SellerUseCase{repository: repository}
}

func (u *SellerUseCase) CreateSeller(sellerInput dto.SellerInputDTO) (dto.SellerOutputDTO, error) {
	seller, err := entity.NewSeller(
		sellerInput.Name,
		sellerInput.Email,
		sellerInput.Phone,
	)

	if err != nil {
		return dto.SellerOutputDTO{}, err
	}

	err = u.repository.CreateSeller(seller)
	if err != nil {
		return dto.SellerOutputDTO{}, err
	}

	sellerOutput := converter.SellerToOutputDTO(seller)

	return sellerOutput, nil
}
