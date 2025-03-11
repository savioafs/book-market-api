package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/usecase"
	"net/http"
)

type DiscountCouponController struct {
	useCase usecase.DiscountCouponUseCase
}

func NewDiscountCouponController(useCase usecase.DiscountCouponUseCase) *DiscountCouponController {
	return &DiscountCouponController{useCase: useCase}
}

func (ct *DiscountCouponController) CreateDiscountCoupon(c *gin.Context) {
	var input dto.DiscountCouponInputDTO

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input to create discount coupon",
			"error":   err.Error(),
		})

		return
	}

	output, err := ct.useCase.CreateDiscountCoupon(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create discount coupon",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, output)
}

func (ct *DiscountCouponController) GetAllDiscountCoupons(c *gin.Context) {

}

func (ct *DiscountCouponController) GetDiscountCoupon(c *gin.Context) {

}

func (ct *DiscountCouponController) GetActiveDiscountsCoupons(c *gin.Context) {

}

func (ct *DiscountCouponController) DisableDiscountCoupon(c *gin.Context) {

}
