package converter

import (
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
)

func BookToOutputDTO(b *entity.Book) dto.BookOutputDTO {
	return dto.BookOutputDTO{
		ID:          b.ID,
		Title:       b.Title,
		ImageURL:    b.ImageURL,
		Author:      b.Author,
		Publisher:   b.Publisher,
		ISBN:        b.ISBN,
		Price:       b.Price,
		Stock:       b.Stock,
		Category:    b.Category,
		Description: b.Description,
		CreatedAt:   b.CreatedAt,
	}
}

func BookToOutputSaleDTO(book *entity.Book) dto.BookForSaleDTO {
	return dto.BookForSaleDTO{
		Title:  book.Title,
		Author: book.Author,
		Price:  book.Price,
	}
}
