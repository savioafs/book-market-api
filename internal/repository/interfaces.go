package repository

import "github.com/savioafs/book-market/internal/entity"

type BookStorer interface {
	CreateBook(book *entity.Book) error
	GetAllBooks() ([]entity.Book, error)
	GetBookByID(id string) (*entity.Book, error)
	GetBooksByCategory(category string) ([]*entity.Book, error)
	GetBooksByPublishedYear(publishedYear int) ([]*entity.Book, error)
	GetBooksByAuthor(author string) ([]*entity.Book, error)
	UpdateBook(book *entity.Book) error
	UpdateStockBookSale(bookID string, quantity int) error
	UpdateStockBookRenew(bookID string, quantity int) error
	DeleteBook(bookID string) error
}

type DiscountCouponStorer interface {
	CreateDiscountCoupon(coupon *entity.DiscountCoupon) error
	GetAllDiscountCoupons() ([]*entity.DiscountCoupon, error)
	CountUse(id string) error
	GetDiscountCoupon(id string) (*entity.DiscountCoupon, error)
	GetActiveDiscountsCoupons() (*[]entity.DiscountCoupon, error)
	DisableDiscountCoupon(id string) error
}

type SellerStorer interface {
	CreateSeller(seller *entity.Seller) error
	GetAllSellers() ([]*entity.Seller, error)
	GetSellerByID(id string) (*entity.Seller, error)
	GetSellerByEmail(email string) (*entity.Seller, error)
	UpdateSeller(seller *entity.Seller) error
	DeleteSeller(id string) error
}

type SaleStorer interface {
	CreateSale(sale *entity.Sale) error
	GetAllSales() ([]*entity.Sale, error)
	GetSaleByID(id string) (*entity.Sale, error)
	UpdateSale(sale *entity.Sale) error
	DeleteSale(id string) error
}

type ReviewStorer interface {
	CreateReview(review *entity.Review) error
	GetReviewByID(id string) (*entity.Review, error)
}
