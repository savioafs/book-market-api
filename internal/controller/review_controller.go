package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/usecase"
)

type ReviewController struct {
	useCase usecase.ReviewUseCase
}

func NewReviewController(useCase usecase.ReviewUseCase) *ReviewController {
	return &ReviewController{useCase: useCase}
}

func (ct *ReviewController) CreateReview(c *gin.Context) {

}

func (ct *ReviewController) GetReviewByID(c *gin.Context) {

}
