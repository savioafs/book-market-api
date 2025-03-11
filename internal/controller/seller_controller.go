package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/usecase"
	"net/http"
)

type SellerController struct {
	useCase usecase.SellerUseCase
}

func NewSellerController(useCase usecase.SellerUseCase) *SellerController {
	return &SellerController{useCase: useCase}
}

func (ct *SellerController) CreateSeller(c *gin.Context) {
	var input dto.SellerInputDTO

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input for create seller",
			"error":   err.Error(),
		})
		return
	}

	output, err := ct.useCase.CreateSeller(input)
	if errors.Is(err, common.ErrSellerAlreadyExists) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "seller already exists",
			"error":   err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create seller",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, output)
}

func (ct *SellerController) GetAllSellers(c *gin.Context) {}

func (ct *SellerController) GetSellerByID(c *gin.Context) {}

func (ct *SellerController) GetSellerByEmail(c *gin.Context) {}

func (ct *SellerController) UpdateSeller(c *gin.Context) {}
func (ct *SellerController) DeleteSeller(c *gin.Context) {}
