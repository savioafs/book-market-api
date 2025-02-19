package usecase

import (
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

	bookOutput := dto.BookOutputDTO{
		ID:          book.ID,
		Title:       book.Title,
		ImageURL:    book.ImageURL,
		Author:      book.Author,
		Publisher:   book.Publisher,
		ISBN:        book.ISBN,
		Price:       book.Price,
		Stock:       book.Stock,
		Category:    book.Category,
		Description: book.Description,
		CreatedAt:   book.CreatedAt,
	}

	return bookOutput, nil
}

func (u *BookUseCase) GetAllBooks() ([]dto.BookOutputDTO, error) {
	var booksOutput []dto.BookOutputDTO

	books, err := u.repository.GetAllBooks()
	if err != nil {
		return []dto.BookOutputDTO{}, err
	}

	for _, book := range books {
		bookOutput := dto.BookOutputDTO{
			ID:          book.ID,
			Title:       book.Title,
			ImageURL:    book.ImageURL,
			Author:      book.Author,
			Publisher:   book.Publisher,
			ISBN:        book.ISBN,
			Price:       book.Price,
			Stock:       book.Stock,
			Category:    book.Category,
			Description: book.Description,
			CreatedAt:   book.CreatedAt,
		}

		booksOutput = append(booksOutput, bookOutput)
	}

	return booksOutput, nil
}

func (u *BookUseCase) GetBookByID(id string) (dto.BookOutputDTO, error) {
	book, err := u.repository.GetBookByID(id)
	if err != nil {
		return dto.BookOutputDTO{}, err
	}

	bookOutput := dto.BookOutputDTO{
		ID:          book.ID,
		Title:       book.Title,
		ImageURL:    book.ImageURL,
		Author:      book.Author,
		Publisher:   book.Publisher,
		ISBN:        book.ISBN,
		Price:       book.Price,
		Stock:       book.Stock,
		Category:    book.Category,
		Description: book.Description,
		CreatedAt:   book.CreatedAt,
	}

	return bookOutput, nil
}

func (u *BookUseCase) GetBookByCategory(category string) ([]dto.BookOutputDTO, error) {
	var booksOutput []dto.BookOutputDTO

	books, err := u.repository.GetBooksByCategory(category)
	if err != nil {
		return []dto.BookOutputDTO{}, err
	}

	for _, book := range books {
		bookOutput := dto.BookOutputDTO{
			ID:          book.ID,
			Title:       book.Title,
			ImageURL:    book.ImageURL,
			Author:      book.Author,
			Publisher:   book.Publisher,
			ISBN:        book.ISBN,
			Price:       book.Price,
			Stock:       book.Stock,
			Category:    book.Category,
			Description: book.Description,
			CreatedAt:   book.CreatedAt,
		}

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

	for _, book := range books {
		bookOutput := dto.BookOutputDTO{
			ID:          book.ID,
			Title:       book.Title,
			ImageURL:    book.ImageURL,
			Author:      book.Author,
			Publisher:   book.Publisher,
			ISBN:        book.ISBN,
			Price:       book.Price,
			Stock:       book.Stock,
			Category:    book.Category,
			Description: book.Description,
			CreatedAt:   book.CreatedAt,
		}

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

	for _, book := range books {
		bookOutput := dto.BookOutputDTO{
			ID:          book.ID,
			Title:       book.Title,
			ImageURL:    book.ImageURL,
			Author:      book.Author,
			Publisher:   book.Publisher,
			ISBN:        book.ISBN,
			Price:       book.Price,
			Stock:       book.Stock,
			Category:    book.Category,
			Description: book.Description,
			CreatedAt:   book.CreatedAt,
		}

		booksOutput = append(booksOutput, bookOutput)
	}

	return booksOutput, nil
}

func (u *BookUseCase) UpdateBook(bookInput dto.BookInputDTO) error {
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
		return err
	}

	err = u.repository.UpdateBook(&book)
	if err != nil {
		return err
	}

	return nil
}

func (u *BookUseCase) UpdateStockBookSale(bookID string, quantity int) error {

	err := u.repository.UpdateStockBookSale(bookID, quantity)
	if err != nil {
		return err
	}

	return nil
}

func (u *BookUseCase) UpdateStockBookRenew(bookID string, quantity int) error {

	err := u.repository.UpdateStockBookRenew(bookID, quantity)
	if err != nil {
		return err
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
