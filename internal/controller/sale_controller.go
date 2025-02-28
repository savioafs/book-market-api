package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/usecase"
	"net/http"
)

type SaleController struct {
	useCase usecase.SaleUseCase
}

func NewSaleController(useCase usecase.SaleUseCase) *SaleController {
	return &SaleController{useCase: useCase}
}

func (ct *SaleController) CreateSale(c *gin.Context) {
	var sale dto.SaleInputDTO

	err := c.BindJSON(&sale)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}

	outPut, err := ct.useCase.CreateSale(sale)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot register sale",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, outPut)
}
func (ct *SaleController) GetAllSales(c *gin.Context) {}
func (ct *SaleController) GetSaleByID(c *gin.Context) {}
func (ct *SaleController) UpdateSale(c *gin.Context)  {}
func (ct *SaleController) DeleteSale(c *gin.Context)  {}
