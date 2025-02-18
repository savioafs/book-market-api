package entity

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var sellertInputTest = Seller{
	Name:  "Marcos",
	Email: "marcos@marcos.net",
	Phone: "27998373245",
}

var bookInputTest = Book{
	Title:         "Title One",
	Author:        "Author One",
	Publisher:     "Publisher One",
	ISBN:          "123456789",
	Price:         110.99,
	Stock:         15,
	Category:      "Category One",
	PublishedYear: 2019,
	Description:   "Description One",
}

var booksInputTest = []Book{
	bookInputTest,
	bookInputTest,
	bookInputTest,
	bookInputTest,
}

var discountInputTest = DiscountCoupon{
	Code:               "february10",
	DiscountPercentage: 10,
	ExpirationDate:     time.Date(2025, time.April, 29, 1, 2, 3, 4, time.UTC),
	UsageLimit:         10,
	UsedCount:          0,
	CreatedAt:          time.Now(),
}

var saleInputTest = Sale{
	BuyerName: "Juliano",
	Quantity:  5,
}

var reviewInputTest = Review{
	Rating:  4.5,
	Comment: "Best seller, very good",
}

func Test_Models(t *testing.T) {
	as := assert.New(t)

	book, err := NewBook(
		bookInputTest.Title,
		bookInputTest.Author,
		bookInputTest.Publisher,
		bookInputTest.ISBN,
		bookInputTest.Category,
		bookInputTest.Description,
		bookInputTest.Price,
		bookInputTest.Stock,
		bookInputTest.PublishedYear,
	)
	as.Nil(err)
	as.NotNil(book)
	as.NotNil(book.ID)
	as.NotNil(book.CreatedAt)
	err = book.Validate()
	as.Nil(err)

	discountCoupon, err := NewDiscountCoupon(
		discountInputTest.Code,
		discountInputTest.DiscountPercentage,
		discountInputTest.ExpirationDate,
		discountInputTest.UsageLimit,
		discountInputTest.UsedCount,
	)
	as.Nil(err)
	as.NotNil(discountCoupon.ID)
	as.NotNil(discountCoupon.Code)
	as.NotNil(discountCoupon.CreatedAt)

	seller, err := NewSeller(
		sellertInputTest.Name,
		sellertInputTest.Email,
		sellertInputTest.Phone,
	)
	as.Nil(err)
	as.NotNil(seller.ID)
	as.NotNil(seller.Code)
	as.NotNil(seller.CreatedAt)

	sale, err := NewSale(booksInputTest, *seller, saleInputTest.BuyerName, *discountCoupon)
	as.Nil(err)
	as.NotNil(sale.ID)
	as.NotNil(sale.Code)
	as.NotNil(sale.CreatedAt)
	as.Equal(sale.Quantity, len(booksInputTest))
	as.Equal(sale.TotalPrice, 443.96)

	review, err := NewReview(*sale, reviewInputTest.Rating, reviewInputTest.Comment)
	as.Nil(err)
	as.NotNil(review.ID)
	as.NotNil(review.CreatedAt)

	fmt.Println(review)

}
