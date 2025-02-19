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

func (ct *DiscountCouponController) CreateDiscountCoupon(c *gin.Context) {

}

func (ct *DiscountCouponController) GetAllDiscountCoupons(c *gin.Context) {

}

func (ct *DiscountCouponController) GetDiscountCoupon(c *gin.Context) {

}

func (ct *DiscountCouponController) GetActiveDiscountsCoupons(c *gin.Context) {

}

func (ct *DiscountCouponController) DisableDiscountCoupon(c *gin.Context) {

}
