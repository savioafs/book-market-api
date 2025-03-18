package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/usecase"
	"net/http"
)

type ReviewController struct {
	useCase usecase.ReviewUseCase
}

func NewReviewController(useCase usecase.ReviewUseCase) *ReviewController {
	return &ReviewController{useCase: useCase}
}

func (ct *ReviewController) CreateReview(c *gin.Context) {
	var input dto.ReviewInputDTO

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}

	reviewOutput, err := ct.useCase.CreateReview(input)
	if errors.Is(err, common.ErrSaleAlreadyReviewed) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "sale already review",
			"error":   err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail to create review",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reviewOutput)
}

func (ct *ReviewController) GetReviewByID(c *gin.Context) {

}
