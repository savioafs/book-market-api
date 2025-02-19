package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/controller"
	"github.com/savioafs/book-market/internal/repository"
	"github.com/savioafs/book-market/internal/usecase"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	// book
	bookRepositoryGorm := repository.NewBookRepositoryGorm(db)
	bookUseCase := usecase.NewBookUseCase(bookRepositoryGorm)
	bookController := controller.NewBookController(*bookUseCase)

	book := router.Group("/book")
	{
		book.POST("/", bookController.CreateBook)
		book.GET("/all", bookController.GetAllBooks)
	}

	// discount coupons
	discountCouponRepositoryGorm := repository.NewDiscountCouponRepositoryGorm(db)
	// discountCouponUseCase
	// discountCouponController
	discountCoupon := router.Group("/discount-coupon")
	{
		discountCoupon.POST("/")
	}

	// review
	reviewRepositoryGorm := repository.NewReviewRepositoryGorm(db)
	// reviewUseCase
	// reviewController
	review := router.Group("/discount-coupon")
	{
		review.POST("/")
	}

	// sale
	saleRepositoryGorm := repository.NewSaleRepositoryGorm(db)
	// saleUseCase
	// saleController
	sale := router.Group("/discount-coupon")
	{
		sale.POST("/")
	}

	// seller
	sellerRepositoryGorm := repository.NewSellerRepositoryGorm(db)
	// sellerUseCase
	// sellerController
	seller := router.Group("/discount-coupon")
	{
		seller.POST("/")
	}

	fmt.Println(bookUseCase, discountCouponRepositoryGorm, reviewRepositoryGorm, saleRepositoryGorm, sellerRepositoryGorm)

}
