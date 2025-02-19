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

func (ct *SaleController) CreateSale(c *gin.Context)  {}
func (ct *SaleController) GetAllSales(c *gin.Context) {}
func (ct *SaleController) GetSaleByID(c *gin.Context) {}
func (ct *SaleController) UpdateSale(c *gin.Context)  {}
func (ct *SaleController) DeleteSale(c *gin.Context)  {}
