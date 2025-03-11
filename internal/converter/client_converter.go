package converter

import (
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
)

func ConvertClientOutput(client *entity.Client) dto.ClientOutputDTO {
	return dto.ClientOutputDTO{
		ID:    client.ID,
		Name:  client.Name,
		Email: client.Email,
		Phone: client.Phone,
	}
}

func ConvertClientSaleOutput(client *entity.Client) dto.ClientOutputForSaleDTO {
	return dto.ClientOutputForSaleDTO{
		Name:  client.Name,
		Email: client.Email,
		Phone: client.Phone,
	}
}
