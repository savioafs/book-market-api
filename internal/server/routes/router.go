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
	}

	// discount coupons
	discountCouponRepositoryGorm := repository.NewDiscountCouponRepositoryGorm(db)

	// review
	reviewRepositoryGorm := repository.NewReviewRepositoryGorm(db)

	// sale
	saleRepositoryGorm := repository.NewSaleRepositoryGorm(db)

	// seller
	sellerRepositoryGorm := repository.NewSellerRepositoryGorm(db)

	fmt.Println(bookUseCase, discountCouponRepositoryGorm, reviewRepositoryGorm, saleRepositoryGorm, sellerRepositoryGorm)

}
