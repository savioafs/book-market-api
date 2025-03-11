package usecase

import (
	"fmt"
	"github.com/savioafs/book-market/internal/actions"
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/converter"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
	"github.com/savioafs/book-market/internal/repository"
)

type BookUseCase struct {
	repository repository.BookStorer
}

func NewBookUseCase(repository repository.BookStorer) *BookUseCase {
	return &BookUseCase{repository: repository}
}

func (u *BookUseCase) CreateBook(bookInput dto.BookInputDTO) (dto.BookOutputDTO, error) {
	bookExists, err := u.repository.ExistsBook(bookInput.Title, bookInput.ImageURL, bookInput.ISBN)
	if err != nil {
		return dto.BookOutputDTO{}, err
	}

	if bookExists {
		return dto.BookOutputDTO{}, common.BookAlreadyExists
	}

	book, err := entity.NewBook(
		bookInput.Title,
		bookInput.ImageURL,
		bookInput.Author,
		bookInput.Publisher,
		bookInput.ISBN,
		bookInput.Category,
		bookInput.Description,
		bookInput.Price,
		bookInput.Stock,
		bookInput.PublishedYear,
	)
	if err != nil {
		return dto.BookOutputDTO{}, err
	}

	err = u.repository.CreateBook(&book)
	if err != nil {
		return dto.BookOutputDTO{}, err
	}

	bookOutput := converter.BookToOutputDTO(&book)

	return bookOutput, nil
}

func (u *BookUseCase) GetAllBooks() ([]dto.BookOutputDTO, error) {
	var booksOutput []dto.BookOutputDTO

	books, err := u.repository.GetAllBooks()
	if err != nil {
		return []dto.BookOutputDTO{}, err
	}

	for _, book := range books {
		bookOutput := converter.BookToOutputDTO(&book)
		booksOutput = append(booksOutput, bookOutput)
	}

	return booksOutput, nil
}

func (u *BookUseCase) GetBookByID(id string) (dto.BookOutputDTO, error) {
	book, err := u.repository.GetBookByID(id)

	if err != nil {
		return dto.BookOutputDTO{}, err
	}

	bookOutput := converter.BookToOutputDTO(book)

	return bookOutput, nil
}

func (u *BookUseCase) GetBookByCategory(category string) ([]dto.BookOutputDTO, error) {
	var booksOutput []dto.BookOutputDTO

	books, err := u.repository.GetBooksByCategory(category)
	if err != nil {
		return []dto.BookOutputDTO{}, err
	}

	if len(books) == 0 {
		return []dto.BookOutputDTO{}, common.ErrCategoryNotFound
	}

	for _, book := range books {
		bookOutput := converter.BookToOutputDTO(book)
		booksOutput = append(booksOutput, bookOutput)
	}

	return booksOutput, nil
}

func (u *BookUseCase) GetBooksByPublishedYear(publishedYear int) ([]dto.BookOutputDTO, error) {
	var booksOutput []dto.BookOutputDTO

	books, err := u.repository.GetBooksByPublishedYear(publishedYear)
	if err != nil {
		return []dto.BookOutputDTO{}, err
	}

	if len(books) == 0 {
		return []dto.BookOutputDTO{}, common.ErrPublishedYearNotFound
	}

	for _, book := range books {
		bookOutput := converter.BookToOutputDTO(book)
		booksOutput = append(booksOutput, bookOutput)
	}

	return booksOutput, nil
}

func (u *BookUseCase) GetBooksByAuthor(author string) ([]dto.BookOutputDTO, error) {
	var booksOutput []dto.BookOutputDTO

	books, err := u.repository.GetBooksByAuthor(author)
	if err != nil {
		return []dto.BookOutputDTO{}, err
	}

	if len(books) == 0 {
		return []dto.BookOutputDTO{}, common.ErrAuthorNotFound
	}

	for _, book := range books {
		bookOutput := converter.BookToOutputDTO(book)
		booksOutput = append(booksOutput, bookOutput)
	}

	return booksOutput, nil
}

func (u *BookUseCase) UpdateBook(bookID string, bookInput dto.BookInputDTO) error {

	book, err := u.repository.GetBookByID(bookID)
	if err != nil {
		return err
	}

	if bookInput.Title != "" {
		book.Title = bookInput.Title
	}
	if bookInput.ImageURL != "" {
		book.ImageURL = bookInput.ImageURL
	}
	if bookInput.Author != "" {
		book.Author = bookInput.Author
	}
	if bookInput.Publisher != "" {
		book.Publisher = bookInput.Publisher
	}
	if bookInput.ISBN != "" {
		book.ISBN = bookInput.ISBN
	}
	if bookInput.Category != "" {
		book.Category = bookInput.Category
	}
	if bookInput.Description != "" {
		book.Description = bookInput.Description
	}
	if bookInput.Price != 0 {
		book.Price = bookInput.Price
	}
	if bookInput.Stock != 0 {
		book.Stock = bookInput.Stock
	}
	if bookInput.PublishedYear != 0 {
		book.PublishedYear = bookInput.PublishedYear
	}

	err = u.repository.UpdateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (u *BookUseCase) UpdateStockBook(bookID string, quantity int, action actions.UpdateBookStockAction) error {
	var err error

	if !action.IsValid() {
		return fmt.Errorf("invalid action for - %s", action)
	}

	switch action {
	case actions.BookRenew:
		err = u.repository.UpdateStockBookRenew(bookID, quantity)
		if err != nil {
			return err
		}
	case actions.BookSale:
		err = u.repository.UpdateStockBookSale(bookID, quantity)
		if err != nil {
			return err
		}

	}

	return nil
}

func (u *BookUseCase) DeleteBook(bookID string) error {

	err := u.repository.DeleteBook(bookID)
	if err != nil {
		return err
	}

	return nil
}
