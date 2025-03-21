package usecase

import (
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/converter"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
	"github.com/savioafs/book-market/internal/repository"
)

// TODO: generate metrics of reviews in seller profile

type SellerUseCase struct {
	repository repository.SellerStorer
}

func NewSellerUseCase(repository repository.SellerStorer) *SellerUseCase {
	return &SellerUseCase{repository: repository}
}

func (u *SellerUseCase) CreateSeller(sellerInput dto.SellerInputDTO) (dto.SellerOutputDTO, error) {
	sellerExists, err := u.repository.ExistsSeller(sellerInput.Name, sellerInput.Email, sellerInput.Phone)
	if err != nil {
		return dto.SellerOutputDTO{}, nil
	}

	if sellerExists {
		return dto.SellerOutputDTO{}, common.ErrSellerAlreadyExists
	}

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
