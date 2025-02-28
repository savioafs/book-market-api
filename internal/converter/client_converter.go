package converter

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
