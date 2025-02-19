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
	discountCouponUseCase := usecase.NewDiscountCouponUseCase(discountCouponRepositoryGorm)
	discountCouponController := controller.NewDiscountCouponController(*discountCouponUseCase)
	discountCoupon := router.Group("/discount-coupon")
	{
		discountCoupon.POST("/", discountCouponController.CreateDiscountCoupon)
	}

	// review
	reviewRepositoryGorm := repository.NewReviewRepositoryGorm(db)
	reviewUseCase := usecase.NewReviewUseCase(reviewRepositoryGorm)
	reviewController := controller.NewReviewController(*reviewUseCase)
	review := router.Group("/review")
	{
		review.POST("/")
	}

	// sale
	saleRepositoryGorm := repository.NewSaleRepositoryGorm(db)
	saleUseCase := usecase.NewSaleUseCase(saleRepositoryGorm)
	saleController := controller.NewSaleController(*saleUseCase)
	sale := router.Group("/sale")
	{
		sale.POST("/")
	}

	// seller
	sellerRepositoryGorm := repository.NewSellerRepositoryGorm(db)
	sellerUseCase := usecase.NewSellerUseCase(sellerRepositoryGorm)
	sellerController := controller.NewSellerController(*sellerUseCase)
	seller := router.Group("/seller")
	{
		seller.POST("/")
	}

	fmt.Println(bookUseCase, discountCouponRepositoryGorm, reviewRepositoryGorm, saleRepositoryGorm, sellerRepositoryGorm)

}
