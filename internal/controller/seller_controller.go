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

func (ct *SellerController) CreateSeller(c *gin.Context) {}

func (ct *SellerController) GetAllSellers(c *gin.Context) {}

func (ct *SellerController) GetSellerByID(c *gin.Context) {}

func (ct *SellerController) GetSellerByEmail(c *gin.Context) {}

func (ct *SellerController) UpdateSeller(c *gin.Context) {}
func (ct *SellerController) DeleteSeller(c *gin.Context) {}
