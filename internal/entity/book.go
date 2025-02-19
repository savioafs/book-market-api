package entity

import (
	"errors"
	"github.com/savioafs/book-market/internal/utils"
	"time"
)

var (
	ErrTitleIsRequired         = errors.New("title is required")
	ErrImageIsRequired         = errors.New("image is required")
	ErrAuthorIsRequired        = errors.New("author is required")
	ErrPublisherIsRequired     = errors.New("publisher is required")
	ErrIsbnIsRequired          = errors.New("isbn is required")
	ErrCategoryIsRequired      = errors.New("category is required")
	ErrDescriptionIsRequired   = errors.New("description is required")
	ErrPriceIsRequired         = errors.New("price is required")
	ErrPublishedYearIsRequired = errors.New("published year is required")
)

type Book struct {
	ID            string    `json:"id" gorm:"primary_key"`
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
		return ErrTitleIsRequired
	}

	if b.ImageURL == "" {
		return ErrImageIsRequired
	}

	if b.Author == "" {
		return ErrAuthorIsRequired
	}

	if b.Publisher == "" {
		return ErrPublisherIsRequired
	}

	if b.ISBN == "" {
		return ErrIsbnIsRequired
	}

	if b.Category == "" {
		return ErrCategoryIsRequired
	}

	if b.Price == 0.0 {
		return ErrPriceIsRequired
	}

	if b.Description == "" {
		return ErrDescriptionIsRequired
	}

	if b.PublishedYear == 0 {
		return ErrPublishedYearIsRequired
	}

	return nil
}
