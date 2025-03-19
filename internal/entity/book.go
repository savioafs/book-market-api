package entity

import (
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/utils"
	"time"
)

type Book struct {
	ID            string    `json:"id" `
	Title         string    `json:"title"`
	ImageURL      string    `json:"image_url"`
	Author        string    `json:"author"`
	Publisher     string    `json:"publisher"`
	ISBN          string    `json:"isbn"`
	Price         float64   `json:"price"`
	Stock         int       `json:"stock"`
	Category      string    `json:"category"`
	PublishedYear int       `json:"published_year"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func NewBook(title, imgURL, author, publisher, isbn, category, description string, price float64, stock, publishedYear int) (Book, error) {
	id, err := utils.NewID()
	if err != nil {
		return Book{}, err
	}

	book := Book{
		ID:            id,
		Title:         title,
		ImageURL:      imgURL,
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

	if err := book.Validate(); err != nil {
		return Book{}, err
	}

	return book, nil
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return common.ErrTitleIsRequired
	}

	if b.ImageURL == "" {
		return common.ErrImageIsRequired
	}

	if b.Author == "" {
		return common.ErrAuthorIsRequired
	}

	if b.Publisher == "" {
		return common.ErrPublisherIsRequired
	}

	if b.ISBN == "" {
		return common.ErrIsbnIsRequired
	}

	if b.Category == "" {
		return common.ErrCategoryIsRequired
	}

	if b.Price == 0.0 {
		return common.ErrPriceIsRequired
	}

	if b.Description == "" {
		return common.ErrDescriptionIsRequired
	}

	if b.PublishedYear == 0 {
		return common.ErrPublishedYearIsRequired
	}

	return nil
}
