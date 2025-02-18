package entity

import (
	"github.com/savioafs/book-market/utils"
	"time"
)

type Book struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Publisher     string    `json:"publisher"`
	ISBN          string    `json:"isbn"`
	Price         float64   `json:"price"`
	Stock         int       `json:"stock"`
	Category      string    `json:"category"`
	PublishedYear int       `json:"published_year"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewBook(title, author, publisher, isbn, category, description string, price float64, stock, publishedYear int) (*Book, error) {
	id, err := utils.NewID()
	if err != nil {
		return nil, err
	}

	book := &Book{
		ID:            id,
		Title:         title,
		Author:        author,
		Publisher:     publisher,
		ISBN:          isbn,
		Price:         price,
		Stock:         stock,
		Category:      category,
		PublishedYear: publishedYear,
		Description:   description,
		CreatedAt:     time.Now(),
	}

	return book, nil
}
