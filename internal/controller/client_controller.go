package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/usecase"
	"net/http"
)

type ClientController struct {
	useCase usecase.ClientUseCase
}

func NewClientController(useCase usecase.ClientUseCase) *ClientController {
	return &ClientController{useCase: useCase}
}

func (u *ClientController) CreateClient(c *gin.Context) {
	var input dto.ClientInputDTO

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})

		return
	}

	output, err := u.useCase.CreateClient(input)
	if errors.Is(err, common.ErrClientAlreadyExists) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email or phone already registred",
			"error":   err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}
