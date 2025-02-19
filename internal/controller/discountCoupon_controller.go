package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/usecase"
)

type DiscountCouponController struct {
	useCase usecase.DiscountCouponUseCase
}

func NewDiscountCouponController(useCase usecase.DiscountCouponUseCase) *DiscountCouponController {
	return &DiscountCouponController{useCase: useCase}
}

func (d *DiscountCouponController) CreateDiscountCoupon(c *gin.Context) {

}

func (d *DiscountCouponController) GetAllDiscountCoupons(c *gin.Context) {

}

func (d *DiscountCouponController) GetDiscountCoupon(c *gin.Context) {

}

func (d *DiscountCouponController) GetActiveDiscountsCoupons(c *gin.Context) {

}

func (d *DiscountCouponController) DisableDiscountCoupon(c *gin.Context) {

}
