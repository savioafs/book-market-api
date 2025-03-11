package usecase

import (
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/converter"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
	"github.com/savioafs/book-market/internal/repository"
)

type ClientUseCase struct {
	repository repository.ClientStorer
}

func NewClientUseCase(repository repository.ClientStorer) *ClientUseCase {
	return &ClientUseCase{repository: repository}
}

func (u *ClientUseCase) CreateClient(clientInput dto.ClientInputDTO) (dto.ClientOutputDTO, error) {
	clientExists, err := u.repository.ExistsClient(clientInput.Name, clientInput.Email, clientInput.Phone)
	if err != nil {
		return dto.ClientOutputDTO{}, nil
	}

	if clientExists {
		return dto.ClientOutputDTO{}, common.ErrClientAlreadyExists
	}

	client, err := entity.NewClient(
		clientInput.Name,
		clientInput.Email,
		clientInput.Phone,
	)
	if err != nil {
		return dto.ClientOutputDTO{}, err
	}

	err = u.repository.CreateClient(&client)
	if err != nil {
		return dto.ClientOutputDTO{}, err
	}

	clientOutput := converter.ConvertClientOutput(&client)

	return clientOutput, nil
}
