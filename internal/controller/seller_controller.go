package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/usecase"
)

type SellerController struct {
	useCase usecase.SellerUseCase
}

func NewSellerController(useCase usecase.SellerUseCase) *SellerController {
	return &SellerController{useCase: useCase}
}

func (sc *SellerController) CreateSeller(c *gin.Context) {}

func (sc *SellerController) GetAllSellers(c *gin.Context) {}

func (sc *SellerController) GetSellerByID(c *gin.Context) {}

func (sc *SellerController) GetSellerByEmail(c *gin.Context) {}

func (sc *SellerController) UpdateSeller(c *gin.Context) {}
func (sc *SellerController) DeleteSeller(c *gin.Context) {}
