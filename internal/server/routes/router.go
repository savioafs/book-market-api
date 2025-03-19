package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/controller"
	"github.com/savioafs/book-market/internal/repository"
	"github.com/savioafs/book-market/internal/usecase"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {

	// book
	bookRepositoryGorm := repository.NewBookRepositoryGorm(db)
	bookUseCase := usecase.NewBookUseCase(bookRepositoryGorm)
	bookController := controller.NewBookController(*bookUseCase)
	book := router.Group("/book")
	{
		book.POST("/", bookController.CreateBook)
		book.GET("/all", bookController.GetAllBooks)
		book.GET("/:id", bookController.GetBookByID)
		book.GET("/category/:category", bookController.GetBooksByCategory)
		book.GET("/published-year/:published_year", bookController.GetBooksByPublishedYear)
		book.GET("/author/:author", bookController.GetBooksByAuthor)
		book.PUT("/:id", bookController.UpdateBook)
		book.PATCH("/update-renew/:id", bookController.UpdateStockBookRenew)
		book.DELETE("/:id", bookController.DeleteBook)
	}

	// discount coupons
	discountCouponRepositoryGorm := repository.NewDiscountCouponRepositoryGorm(db)
	discountCouponUseCase := usecase.NewDiscountCouponUseCase(discountCouponRepositoryGorm)
	discountCouponController := controller.NewDiscountCouponController(*discountCouponUseCase)
	discountCoupon := router.Group("/discount-coupon")
	{
		discountCoupon.POST("/", discountCouponController.CreateDiscountCoupon)
		discountCoupon.GET("/", discountCouponController.GetAllDiscountCoupons)
		discountCoupon.GET("/:id", discountCouponController.GetDiscountCoupon)
		discountCoupon.GET("/active", discountCouponController.GetActiveDiscountsCoupons)
		discountCoupon.PATCH("/disable/:id", discountCouponController.DisableDiscountCoupon)
	}

	// seller
	sellerRepositoryGorm := repository.NewSellerRepositoryGorm(db)
	sellerUseCase := usecase.NewSellerUseCase(sellerRepositoryGorm)
	sellerController := controller.NewSellerController(*sellerUseCase)
	seller := router.Group("/seller")
	{
		seller.POST("/", sellerController.CreateSeller)
		seller.GET("/", sellerController.GetAllSellers)
		seller.GET("/:id", sellerController.GetSellerByID)
		seller.GET("/email/:email", sellerController.GetSellerByEmail)
		seller.PUT("/:id", sellerController.UpdateSeller)
		seller.DELETE("/:id", sellerController.DeleteSeller)
	}

	// client
	clientRepositoryGorm := repository.NewClientRepositoryGorm(db)
	clientUseCase := usecase.NewClientUseCase(clientRepositoryGorm)
	clientController := controller.NewClientController(*clientUseCase)
	client := router.Group("/client")
	{
		client.POST("/", clientController.CreateClient)
	}

	// sale
	saleRepositoryGorm := repository.NewSaleRepositoryGorm(db)
	saleUseCase := usecase.NewSaleUseCase(saleRepositoryGorm, bookRepositoryGorm, sellerRepositoryGorm, discountCouponRepositoryGorm, clientRepositoryGorm)
	saleController := controller.NewSaleController(*saleUseCase)
	sale := router.Group("/sale")
	{
		sale.POST("/", saleController.CreateSale)
		sale.POST("/all", saleController.GetAllSales)
		sale.POST("/:id", saleController.GetSaleByID)
		sale.PUT("/:id", saleController.UpdateSale)
		sale.DELETE("/:id", saleController.DeleteSale)
	}

	// review
	reviewRepositoryGorm := repository.NewReviewRepositoryGorm(db)
	reviewUseCase := usecase.NewReviewUseCase(reviewRepositoryGorm, saleRepositoryGorm)
	reviewController := controller.NewReviewController(*reviewUseCase)
	review := router.Group("/review")
	{
		review.POST("/", reviewController.CreateReview)
		review.GET("/", reviewController.GetReviewByID)
	}
}
