package repository

import (
	"github.com/google/uuid"
	"testing"

	"github.com/savioafs/book-market/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&entity.Book{})
	return db
}

var bookToTest = entity.Book{
	Title:         "Golang Testing",
	ImageURL:      "https://google.com",
	Author:        "John Doe",
	Publisher:     "TechBooks",
	ISBN:          "123456789",
	Price:         49.99,
	Stock:         10,
	Category:      "Programming",
	PublishedYear: 2023,
	Description:   "A book about Golang Testing",
}

func TestCreateBook(t *testing.T) {
	as := assert.New(t)

	db := setupTestDB()
	repo := NewBookRepositoryPG(db)

	book, err := entity.NewBook(
		bookToTest.Title,
		bookToTest.ImageURL,
		bookToTest.Author,
		bookToTest.Publisher,
		bookToTest.ISBN,
		bookToTest.Category,
		bookToTest.Description,
		bookToTest.Price,
		bookToTest.Stock,
		bookToTest.PublishedYear,
	)

	err = repo.CreateBook(book)
	as.NoError(err)

	var result entity.Book
	db.First(&result, "id = ?", book.ID)

	as.Equal(book.Title, result.Title)
	as.Equal(book.Author, result.Author)
	as.Equal(book.Stock, result.Stock)
}

func TestGetBookByID(t *testing.T) {
	as := assert.New(t)

	db := setupTestDB()
	repo := NewBookRepositoryPG(db)

	book, err := entity.NewBook(
		bookToTest.Title,
		bookToTest.ImageURL,
		bookToTest.Author,
		bookToTest.Publisher,
		bookToTest.ISBN,
		bookToTest.Category,
		bookToTest.Description,
		bookToTest.Price,
		bookToTest.Stock,
		bookToTest.PublishedYear,
	)

	err = repo.CreateBook(book)
	as.NoError(err)

	result, err := repo.GetBookByID(book.ID)
	as.NoError(err)
	as.Equal(book.Title, result.Title)
}

func TestGetBooksByCategory(t *testing.T) {
	as := assert.New(t)

	db := setupTestDB()
	repo := NewBookRepositoryPG(db)

	var books []*entity.Book

	for i := 0; i < 10; i++ {
		book, err := entity.NewBook(
			bookToTest.Title,
			bookToTest.ImageURL,
			bookToTest.Author,
			bookToTest.Publisher,
			bookToTest.ISBN,
			bookToTest.Category,
			bookToTest.Description,
			bookToTest.Price,
			bookToTest.Stock,
			bookToTest.PublishedYear,
		)
		as.Nil(err)

		err = repo.CreateBook(book)
		as.NoError(err)

		books = append(books, &book)
	}

	result, err := repo.GetBooksByCategory(bookToTest.Category)
	as.NoError(err)
	as.Equal(books[0].Title, result[0].Title)
	as.Equal(len(books), len(result))
}

func TestGetBooksByPublishedYear(t *testing.T) {
	as := assert.New(t)

	db := setupTestDB()
	repo := NewBookRepositoryPG(db)

	var books []*entity.Book

	for i := 0; i < 10; i++ {
		book, err := entity.NewBook(
			bookToTest.Title,
			bookToTest.ImageURL,
			bookToTest.Author,
			bookToTest.Publisher,
			bookToTest.ISBN,
			bookToTest.Category,
			bookToTest.Description,
			bookToTest.Price,
			bookToTest.Stock,
			bookToTest.PublishedYear,
		)
		as.Nil(err)

		err = repo.CreateBook(book)
		as.NoError(err)

		books = append(books, &book)
	}

	result, err := repo.GetBooksByPublishedYear(bookToTest.PublishedYear)
	as.NoError(err)
	as.Equal(books[0].Title, result[0].Title)
	as.Equal(len(books), len(result))
}

func TestGetBooksByAuthor(t *testing.T) {
	as := assert.New(t)

	db := setupTestDB()
	repo := NewBookRepositoryPG(db)

	var books []*entity.Book

	for i := 0; i < 10; i++ {
		book, err := entity.NewBook(
			bookToTest.Title,
			bookToTest.ImageURL,
			bookToTest.Author,
			bookToTest.Publisher,
			bookToTest.ISBN,
			bookToTest.Category,
			bookToTest.Description,
			bookToTest.Price,
			bookToTest.Stock,
			bookToTest.PublishedYear,
		)
		as.Nil(err)

		err = repo.CreateBook(book)
		as.NoError(err)

		books = append(books, &book)
	}

	result, err := repo.GetBooksByAuthor(bookToTest.Author)
	as.NoError(err)
	as.Equal(books[0].Title, result[0].Title)
	as.Equal(len(books), len(result))
}

func TestUpdateBook(t *testing.T) {
	as := assert.New(t)
	db := setupTestDB()
	repo := NewBookRepositoryPG(db)

	book := entity.Book{
		ID:     uuid.New().String(),
		Title:  "Initial Title",
		Author: "Initial Author",
	}

	err := repo.CreateBook(book)
	as.NoError(err)

	book.Title = "Title Updated"

	err = repo.UpdateBook(&book)
	as.NoError(err)

	res, err := repo.GetBookByID(book.ID)
	as.NoError(err)

	as.Equal("Title Updated", res.Title)
	as.Equal("Initial Author", res.Author)
}

func TestUpdateStockBookSale(t *testing.T) {
	as := assert.New(t)

	db := setupTestDB()
	repo := NewBookRepositoryPG(db)

	book, err := entity.NewBook(
		bookToTest.Title,
		bookToTest.ImageURL,
		bookToTest.Author,
		bookToTest.Publisher,
		bookToTest.ISBN,
		bookToTest.Category,
		bookToTest.Description,
		bookToTest.Price,
		bookToTest.Stock,
		bookToTest.PublishedYear,
	)

	err = repo.CreateBook(book)
	as.NoError(err)

	err = repo.UpdateStockBookSale(book.ID, 2)
	as.NoError(err)

	var result entity.Book
	db.First(&result, "id = ?", book.ID)
	as.Equal(result.Stock, book.Stock-2)
}

func TestUpdateStockBookRenew(t *testing.T) {
	as := assert.New(t)

	db := setupTestDB()
	repo := NewBookRepositoryPG(db)

	book, err := entity.NewBook(
		bookToTest.Title,
		bookToTest.ImageURL,
		bookToTest.Author,
		bookToTest.Publisher,
		bookToTest.ISBN,
		bookToTest.Category,
		bookToTest.Description,
		bookToTest.Price,
		bookToTest.Stock,
		bookToTest.PublishedYear,
	)

	err = repo.CreateBook(book)
	as.NoError(err)

	err = repo.UpdateStockBookRenew(book.ID, 50)
	as.NoError(err)

	var result entity.Book
	db.First(&result, "id = ?", book.ID)
	as.Equal(result.Stock, book.Stock+50)
}

func TestDeleteBook(t *testing.T) {
	as := assert.New(t)
	db := setupTestDB()
	repo := NewBookRepositoryPG(db)

	bookId := uuid.New().String()

	book := entity.Book{
		ID:    bookId,
		Title: "New Title Updated",
	}
	db.Create(&book)

	err := repo.DeleteBook(bookId)
	as.NoError(err)

	var result entity.Book
	err = db.First(&result, "id = ?", bookId).Error
	as.Error(err)
	as.Equal(err, gorm.ErrRecordNotFound)
}
