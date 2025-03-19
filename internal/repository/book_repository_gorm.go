package repository

import (
	"database/sql"
	"github.com/savioafs/book-market/internal/entity"
)

type BookRepositoryGorm struct {
	DB *sql.DB
}

func NewBookRepositoryGorm(db *sql.DB) *BookRepositoryGorm {
	return &BookRepositoryGorm{DB: db}
}

func (r *BookRepositoryGorm) CreateBook(b *entity.Book) error {
	stmt, err := r.DB.Prepare("insert into books(id, title, image_url, author, publisher, isbn, price, stock, category, published_year, description, created_at) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		b.ID,
		b.Title,
		b.ImageURL,
		b.Author,
		b.Publisher,
		b.ISBN,
		b.Price,
		b.Stock,
		b.Category,
		b.PublishedYear,
		b.Description,
		b.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

/*
func (r *BookRepositoryGorm) GetAllBooks() ([]entity.Book, error) {
	var books []entity.Book

	err := r.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepositoryGorm) GetBookByID(id string) (*entity.Book, error) {
	var book *entity.Book

	err := r.DB.First(&book, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.ErrBookNotFound
	}
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepositoryGorm) GetBooksByIDs(ids []string) ([]entity.Book, error) {
	var books []entity.Book

	err := r.DB.Where("id IN (?)", ids).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
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

func (r *BookRepositoryGorm) ExistsBook(title, imageUrl, isbn string) (bool, error) {
	var count int64

	err := r.DB.Model(&entity.Book{}).Where("lower(title) = lower(?) AND image_url = ? AND isbn = ?", title, imageUrl, isbn).Count(&count).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
*/
