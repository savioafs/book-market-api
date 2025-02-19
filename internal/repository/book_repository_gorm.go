package repository

import (
	"errors"
	"github.com/savioafs/book-market/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type BookRepositoryGorm struct {
	DB *gorm.DB
}

func NewBookRepositoryGorm(db *gorm.DB) *BookRepositoryGorm {
	return &BookRepositoryGorm{DB: db}
}

func (r *BookRepositoryGorm) CreateBook(book *entity.Book) error {
	return r.DB.Create(&book).Error
}

func (r *BookRepositoryGorm) GetBookByID(id string) (*entity.Book, error) {
	var book *entity.Book

	err := r.DB.First(&book, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepositoryGorm) GetBooksByCategory(category string) ([]*entity.Book, error) {
	var books []*entity.Book

	err := r.DB.Where("category = ?", category).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepositoryGorm) GetBooksByPublishedYear(publishedYear int) ([]*entity.Book, error) {
	var books []*entity.Book

	err := r.DB.Where("published_year = ?", publishedYear).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepositoryGorm) GetBooksByAuthor(author string) ([]*entity.Book, error) {
	var books []*entity.Book

	err := r.DB.Where("author = ?", author).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepositoryGorm) UpdateBook(book *entity.Book) error {
	err := r.DB.Model(&entity.Book{}).Where("id = ?", book.ID).Updates(map[string]interface{}{
		"title":          book.Title,
		"image_url":      book.ImageURL,
		"author":         book.Author,
		"publisher":      book.Publisher,
		"isbn":           book.ISBN,
		"category":       book.Category,
		"description":    book.Description,
		"price":          book.Price,
		"stock":          book.Stock,
		"published_year": book.PublishedYear,
		"updated_at":     time.Now(),
	}).Error

	return err
}

func (r *BookRepositoryGorm) UpdateStockBookSale(bookID string, quantity int) error {
	var b *entity.Book

	err := r.DB.Clauses(clause.Locking{Strength: "UPDATE"}).First(&b, "id = ?", bookID).Error
	if err != nil {
		return err
	}

	err = r.DB.Model(&b).UpdateColumn("stock", gorm.Expr("stock - ?", quantity)).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepositoryGorm) UpdateStockBookRenew(bookID string, quantity int) error {
	var b *entity.Book

	err := r.DB.Clauses(clause.Locking{Strength: "UPDATE"}).First(&b, "id = ?", bookID).Error
	if err != nil {
		return err
	}

	err = r.DB.Model(&b).UpdateColumn("stock", gorm.Expr("stock + ?", quantity)).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *BookRepositoryGorm) DeleteBook(bookID string) error {
	res := r.DB.Delete(&entity.Book{}, "id = ?", bookID)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}
