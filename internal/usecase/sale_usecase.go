package usecase

import (
	"errors"
	"github.com/savioafs/book-market/internal/converters"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/entity"
	"github.com/savioafs/book-market/internal/repository"
	"log"
)

var (
	ErrSellerNotFound        = errors.New("seller not found")
	ErrInvalidDiscountCoupon = errors.New("invalid discount coupon")
)

type SaleUseCase struct {
	saleRepository           repository.SaleStorer
	bookRepository           repository.BookStorer
	sellerRepository         repository.SellerStorer
	discountCouponRepository repository.DiscountCouponStorer
	clientRepository         repository.ClientStorer
}

func NewSaleUseCase(saleRepository repository.SaleStorer, bookRepository repository.BookStorer, sellerRepository repository.SellerStorer, discountCouponRepository repository.DiscountCouponStorer, clientRepository repository.ClientStorer) *SaleUseCase {
	return &SaleUseCase{
		saleRepository:           saleRepository,
		bookRepository:           bookRepository,
		sellerRepository:         sellerRepository,
		discountCouponRepository: discountCouponRepository,
	}
}

func (u *SaleUseCase) CreateSale(saleInput dto.SaleInputDTO) (dto.SaleOutputDTO, error) {

	books, err := u.bookRepository.GetBooksByIDs(saleInput.BooksIDs)
	if err != nil {
		return dto.SaleOutputDTO{}, err
	}

	seller, err := u.sellerRepository.GetSellerByID(saleInput.SellerID)
	if err != nil {
		return dto.SaleOutputDTO{}, err
	}

	if seller == nil {
		return dto.SaleOutputDTO{}, ErrSellerNotFound
	}
	var discountCoupon *entity.DiscountCoupon
	if saleInput.DiscountCouponID != "" {
		discountCoupon, err = u.discountCouponRepository.GetDiscountCoupon(saleInput.DiscountCouponID)
		if err != nil {
			return dto.SaleOutputDTO{}, err
		}
		if discountCoupon == nil {
			log.Printf("invlaid discount coupon id %s", saleInput.DiscountCouponID)
		}
	}

	client, err := u.clientRepository.GetClientByPhone(saleInput.ClientPhone)
	if err != nil {
		return dto.SaleOutputDTO{}, err
	}

	sale, err := entity.NewSale(
		books,
		*seller,
		*client,
		*discountCoupon,
	)
	if err != nil {
		return dto.SaleOutputDTO{}, err
	}

	discountPercentage, finalPrice, err := sale.ApplyDiscountPercentageAndFinalPrice(discountCoupon)
	if err != nil {
		return dto.SaleOutputDTO{}, err
	}

	err = u.saleRepository.CreateSale(sale)
	if err != nil {
		return dto.SaleOutputDTO{}, err
	}

	booksOutput := make([]dto.BookForSaleDTO, len(books))
	for i, book := range books {
		booksOutput[i] = converters.ConvertBookOutput(&book)
	}

	sellerOutput := converters.ConvertSellerOutput(seller)
	clientOutput := converters.ConvertClientOutput(client)
	saleOutput := converters.ConvertSaleOutput(sale, discountCoupon, booksOutput, sellerOutput, clientOutput, discountPercentage, finalPrice)

	return saleOutput, nil
}
