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

func (r *ReviewController) CreateReview(c *gin.Context) {

}

func (r *ReviewController) GetReviewByID(c *gin.Context) {

}
