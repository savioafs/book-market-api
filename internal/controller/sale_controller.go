package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/usecase"
)

type SaleController struct {
	useCase usecase.SaleUseCase
}

func NewSaleController(useCase usecase.SaleUseCase) *SaleController {
	return &SaleController{useCase: useCase}
}

func (sc *SaleController) CreateSale(c *gin.Context)  {}
func (sc *SaleController) GetAllSales(c *gin.Context) {}
func (sc *SaleController) GetSaleByID(c *gin.Context) {}
func (sc *SaleController) UpdateSale(c *gin.Context)  {}
func (sc *SaleController) DeleteSale(c *gin.Context)  {}
